package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
	xj "github.com/basgys/goxml2json"
	"github.com/labstack/echo/v4"
	"time"
)

type (
	// Description of an existing VM
	vmGet struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Logo     string `json:"logo"`
		TID      string `json:"tid"`
		IPv4     string `json:"ipv4"`
		IPv6     string `json:"ipv6,omitempty"`
		State    string `json:"state"`
		LCMState string `json:"lcmstate"`
		Time     time.Time    `json:"time"`
		Disk     int    `json:"disk"`
	}
	// Array of VMS
	vmsGet struct {
		VMs []vmGet `json:"vms"`
	}
	// Description to create a VM
	vmPost struct {
		ID       int    `json:"tid" form:"tid"`
		Name     string `json:"name" form:"name"`
		VCPU     int    `json:"cpu" form:"cpu"`
		Memory   int    `json:"memory" form:"memory"`
		Password string `json:"password" form:"password"`
		SSHKey   string `json:"sshkey" form:"sshkey"`
		Disk     int    `json:"disk" form:"disk"`
	}
	// Snapshot
	snapshotGet struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Time time.Time    `json:"time"`
	}
	// Snapshots
	snapshotsGet struct {
		Snapshots []snapshotGet `json:"snapshots"`
	}
	// Snapshot name
	snapshotCreatePost struct {
		Name string `json:"name" form:"name"`
	}
	// Actions
	actionConfig struct {
		Type   string `json:"type" form:"type"`
		Name   string `json:"name" form:"name"`
		Memory int    `json:"memory" form:"memory"`
		CPU    int    `json:"cpu" form:"cpu"`
		Disk   int    `json:"disk" form:"disk"`
	}
)

var (
	// Regexes for extracting info from OpenNebula's API response
	templateIDRegex   = regexp.MustCompile(`TEMPLATE_ID=\"(.*)\"`)
	imageRegex        = regexp.MustCompile(`IMAGE=\"(.*)\"`)
	imageunameRegex   = regexp.MustCompile(`IMAGE_UNAME=\"(.*)\"`)
	disksizeRegex     = regexp.MustCompile(` SIZE=\"(.*)\"`)
	ipv4Regex         = regexp.MustCompile(`ETH0_IP=\"(.*)\"`)
	ipv6Regex         = regexp.MustCompile(`ETH0_IP6=\"(.*)\"`)
	snapshotRegex     = regexp.MustCompile(`(?s)SNAPSHOT=\[(.*?)\" ]`)
	snapshotNameRegex = regexp.MustCompile(`NAME=\"(.*)\"`)
	snapshotIDRegex   = regexp.MustCompile(`SNAPSHOT_ID=\"(.*)\"`)
	snapshotTimeRegex = regexp.MustCompile(`TIME=\"(.*)\"`)
	csrfRegex         = regexp.MustCompile(`(?s)var csrftoken = '(.*)';`)
)

// getVMs returns the general details[id, name, ip, lcm state] of all the VMs accessible by user.
func getVMs(c echo.Context) error {
	one := c.Get("one").(*goca.Controller)
	vms, err := one.VMs().Info(-1)
	if err == nil {
		vmslist := make([]vmGet, len(vms.VMs))
		for i := 0; i < len(vms.VMs); i++ {
			vm, _ := one.VM(vms.VMs[i].ID).Info(false)
			vmslist[i].ID = vm.ID
			vmslist[i].Name = vm.Name
			vmslist[i].TID = templateIDRegex.FindStringSubmatch(vm.Template.String())[1]
			vmslist[i].Logo = xonglURL + "/" + logoRegex.FindStringSubmatch(vm.UserTemplate.String())[1]
			vmslist[i].IPv4 = ipv4Regex.FindStringSubmatch(vm.Template.String())[1]
			vmslist[i].IPv6 = ipv6Regex.FindStringSubmatch(vm.Template.String())[1]
			vmslist[i].State, vmslist[i].LCMState, _ = vm.StateString()
			vmslist[i].Time = time.Unix(int64(vm.STime), 0)
			diskSize := disksizeRegex.FindStringSubmatch(vm.Template.String())[1]
			vmslist[i].Disk, _ = strconv.Atoi(diskSize)
		}
		vmsjson := &vmsGet{
			VMs: vmslist,
		}
		return c.JSON(http.StatusOK, vmsjson)
	}
	return c.JSON(http.StatusBadRequest, errorMsg(err))
}

// Get the details of a VM
func getVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		one := c.Get("one").(*goca.Controller)
		vm, err := one.VM(id).Info(false)
		if err == nil {
			vminfo := &vmGet{}
			vminfo.ID = vm.ID
			vminfo.Name = vm.Name
			vminfo.TID = templateIDRegex.FindStringSubmatch(vm.Template.String())[1]
			vminfo.Logo = xonglURL + "/" + logoRegex.FindStringSubmatch(vm.UserTemplate.String())[1]
			vminfo.IPv4 = ipv4Regex.FindStringSubmatch(vm.Template.String())[1]
			vminfo.IPv6 = ipv6Regex.FindStringSubmatch(vm.Template.String())[1]
			vminfo.State, vminfo.LCMState, _ = vm.StateString()
			vminfo.Time = time.Unix(int64(vm.STime), 0)
			diskSize := disksizeRegex.FindStringSubmatch(vm.Template.String())[1]
			vminfo.Disk, _ = strconv.Atoi(diskSize)

			return c.JSON(http.StatusOK, vminfo)
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// createVM creates the VM using details[tid, name, memory, password/sshkey, vcpus, disk size]
func createVM(c echo.Context) error {
	config := new(vmPost)
	one := c.Get("one").(*goca.Controller)
	if err := c.Bind(config); err == nil {
		if config.ID != 0 {
			if config.Password != "" || config.SSHKey != "" {
				templateString := ""
				if config.Memory != 0 {
					memory := strconv.Itoa(config.Memory)
					templateString = templateString + "MEMORY=" + memory + "\n"
				}
				if config.Disk != 0 {
					size := strconv.Itoa(config.Disk)
					templateinfo, err := one.Template(config.ID).Info(true, false)
					if err == nil {
						image := imageRegex.FindStringSubmatch(templateinfo.Template.String())[1]
						imageuname := imageunameRegex.FindStringSubmatch(templateinfo.Template.String())[1]
						if config.Disk < 122880 {
							templateString = templateString + "DISK=[IMAGE=\"" + image + "\"," + "IMAGE_UNAME=\"" + imageuname + "\"," + "SIZE=\"" + size + "\"]\n"
						} else {
							templateString = templateString + "DISK=[IMAGE=\"" + image + "\"," + "IMAGE_UNAME=\"" + imageuname + "\"," + "SIZE=\"" + "122880" + "\"]\n"
						}
					} else {
						return c.JSON(http.StatusBadRequest, errorMsg(err))
					}
				}
				if config.VCPU != 0 {
					vcpu := strconv.Itoa(config.VCPU)
					templateString = templateString + "VCPU=" + vcpu + "\n"
				}
				if config.Password != "" {
					password := b64.StdEncoding.EncodeToString([]byte(config.Password))
					templateString = templateString + "PASSWORD_BASE64=" + "\"" + password + "\"" + "\n"
				}
				if config.SSHKey != "" {
					templateString = templateString + "SSH_PUBLIC_KEY=" + "\"" + config.SSHKey + "\"" + "\n"
				}
				if config.Name != "" {
					templateString = templateString + "SET_HOSTNAME=" + "\"" + config.Name + "\"" + "\n"
				}
				id, err := one.Template(config.ID).Instantiate(config.Name, false, templateString, false)
				if err == nil {
					return c.JSON(http.StatusCreated, okMsg("%v[%v] is created\n", config.Name, id))
				}
				return c.JSON(http.StatusBadRequest, errorMsg(err))
			}
			return c.JSON(http.StatusBadRequest, errorJSON{"Please provide Password or SSHKey to create VM"})
		}
		return c.JSON(http.StatusBadRequest, errorJSON{"Please provide a template ID to create VM"})
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid data!"})
}

// Delete VM
func deleteVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		one := c.Get("one").(*goca.Controller)
		err := one.VM(id).Action("terminate-hard")
		if err == nil {
			return c.JSON(http.StatusOK, okMsg("VM is deleted"))
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Submit actions such as reboot <--hard>, power off <--hard>, resume, suspend
func actionVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		config := new(actionConfig)
		one := c.Get("one").(*goca.Controller)
		if err := c.Bind(config); err == nil {
			switch config.Type {
			case "reboot", "reboot-hard", "poweroff", "poweroff-hard", "suspend", "resume","stop":
				err = one.VM(id).Action(config.Type)
			case "rename":
				err = one.VM(id).Rename(config.Name)
			case "resize":
				var templateString string
				if config.Memory != 0 {
					templateString = templateString + "MEMORY=" + strconv.Itoa(config.Memory) + "\n"
 				}
				if config.CPU != 0 {
					templateString = templateString + "VCPU=" + strconv.Itoa(config.CPU) + "\n"
 				}
				err = one.VM(id).Resize(templateString, true)
			case "disk-resize":
				err = one.VM(id).Disk(0).Resize(strconv.Itoa(config.Disk))
			default:
				return c.JSON(http.StatusBadRequest, errorJSON{`Invalid action for the VM`})
			}
			if err == nil {
				return c.JSON(http.StatusOK, okMsg(`VM[%d] action '%v' done successfully`, id, config.Type))
			}
			return c.JSON(http.StatusBadRequest, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorJSON{"No action has been posted"})
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Return the monitoring info of VM using its ID
func monitorVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		one := c.Get("one").(*goca.Controller)
		monitorinfo, err := one.VM(id).Monitoring()
		if err == nil {
			monitorxml := strings.NewReader(monitorinfo)
			monitorbytes, _ := xj.Convert(monitorxml)
			var monitorjson interface{}
			json.Unmarshal(monitorbytes.Bytes(), &monitorjson)
			return c.JSON(http.StatusOK, monitorjson)
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Return the list of snapshot
func snapshotlistVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		one := c.Get("one").(*goca.Controller)
		vminfo, err := one.VM(id).Info(false)
		if err == nil {
			snapshots := snapshotRegex.FindAllStringSubmatch(vminfo.Template.String(), -1)
			snapshotGets := make([]snapshotGet, len(snapshots))
			for i := 0; i < len(snapshots); i++ {
				snapshotGets[i].ID, _ = strconv.Atoi(snapshotIDRegex.FindStringSubmatch(snapshots[i][0])[1])
				snapshotGets[i].Name = snapshotNameRegex.FindStringSubmatch(snapshots[i][0])[1]
				t, _   := strconv.ParseInt((snapshotTimeRegex.FindStringSubmatch(snapshots[i][0])[1]), 10, 64)
				snapshotGets[i].Time  = time.Unix(t, 0)
			}
			snapshotsjson := &snapshotsGet{
				Snapshots: snapshotGets,
			}
			return c.JSON(http.StatusOK, snapshotsjson)
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Create a snapshot of the VM
func snapshotcreateVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		config := new(snapshotCreatePost)
		one := c.Get("one").(*goca.Controller)
		if err := c.Bind(config); err == nil {
			err = one.VM(id).SnapshotCreate(config.Name)
			if err == nil {
				return c.JSON(http.StatusCreated, okMsg("VM[%d] SNAPSHOT[%s] has been created.", id, config.Name))
			}
			return c.JSON(http.StatusBadRequest, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Delete the snapshot of the VM
func snapshotdeleteVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		sid, err := strconv.Atoi(c.Param("sid"))
		if err == nil {
			one := c.Get("one").(*goca.Controller)
			err = one.VM(id).SnapshotDelete(sid)
			if err == nil {
				return c.JSON(http.StatusOK, okMsg("VM[%d] SNAPSHOT[%d] has been deleted.", id, sid))
			}
			return c.JSON(http.StatusBadRequest, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorJSON{"Invalid SNAPSHOT ID[id should be an integer]"})
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Restore the snapshot of the VM
func snapshotrestoreVM(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		sid, err := strconv.Atoi(c.Param("sid"))
		if err == nil {
			one := c.Get("one").(*goca.Controller)
			err = one.VM(id).SnapshotRevert(sid)
			if err == nil {
				return c.JSON(http.StatusOK, okMsg("VM[%d] SNAPSHOT[%d] has been restored.", id, sid))
			}
			return c.JSON(http.StatusBadRequest, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorJSON{"Invalid SNAPSHOT ID[id should be an integer]"})
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}

// Start VNC
func startVNC(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		conf := c.Get("conf").(goca.OneConfig)
		token := strings.Split(conf.Token, ":")
		username, password := token[0], token[1]
		httpclient := &http.Client{}
		loginreq, _ := http.NewRequest("POST", xonglURL+"/login", nil)
		loginreq.SetBasicAuth(username, password)
		loginres, err := httpclient.Do(loginreq)
		if err == nil {
			rootreq, _ := http.NewRequest("GET", xonglURL+"/", nil)
			for _, logincookie := range loginres.Cookies() {
				rootreq.AddCookie(logincookie)
			}
			rootres, err := httpclient.Do(rootreq)
			if err == nil {
				responsedata, _ := ioutil.ReadAll(rootres.Body)
				csrftoken := csrfRegex.FindStringSubmatch(string(responsedata))[1]
				form := url.Values{"csrftoken": {csrftoken}}
				vncreq, _ := http.NewRequest("POST", xonglURL+"/vm/"+strconv.Itoa(id)+"/startvnc", strings.NewReader(form.Encode()))
				for _, logincookie := range loginres.Cookies() {
					vncreq.AddCookie(logincookie)
				}
				for _, rootcookie := range rootres.Cookies() {
					vncreq.AddCookie(rootcookie)
				}
				vncreq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
				vncres, err := httpclient.Do(vncreq)
				if err == nil {
					responsedata, _ = ioutil.ReadAll(vncres.Body)
					return c.JSON(http.StatusOK, responsedata)
				}
				return c.JSON(http.StatusBadRequest, errorMsg(err))
			}
			return c.JSON(http.StatusBadRequest, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid VM ID[id should be an integer]"})
}
