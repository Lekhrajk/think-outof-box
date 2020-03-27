package main

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
	"github.com/labstack/echo/v4"
)

type (
	// Description of an existing template
	templateGet struct {
		ID   int    `json:"id" form:"id"`
		Name string `json:"name" form:"name"`
		Logo string `json:"logo" form:"logo"`
	}
	// Array of a templates
	templatesGet struct {
		Templates []templateGet `json:"templates"`
	}
	// Description of an existing template
	templateDetGet struct {
		ID        int    `json:"id" form:"id"`
		Name      string `json:"name" form:"name"`
		Logo      string `json:"logo" form:"logo"`
		MinVCPU   int    `json:"minVCPU" form:"minVCPU"`
		MaxVCPU   int    `json:"maxVCPU" form:"maxVCPU"`
		MinMEMORY int    `json:"minMEMORY" form:"minMEMORY"`
		MaxMEMORY int    `json:"maxMEMORY" form:"maxMEMORY"`
		MinDisk   int    `json:"minDisk" form:"minDisk"`
		MaxDisk   int    `json:"maxDisk" form:"maxDisk"`
	}
	// Array of a Description templates
	templatesDetGet struct {
		Templates []templateDetGet `json:"templates"`
	}
)

var (
	diskMax     = 122880
	logoRegex   = regexp.MustCompile(`LOGO=\"(.*)\"`)
	vcpuRegex   = regexp.MustCompile(`VCPU=\"O\|range\|\|(\d)\.\.(\d)\|(\d)\"`)
	memoryRegex = regexp.MustCompile(`MEMORY=\"M\|range\|\|(\d+)\.\.(\d+)\|(\d+)\"`)
)

func getTemplates(c echo.Context) error {
	one := c.Get("one").(*goca.Controller)
	templates, err := one.Templates().Info(-1)
	if err == nil {
		templatesjson := make([]templateGet, len(templates.Templates))
		for i := 0; i < len(templates.Templates); i++ {
			templatesjson[i].ID = templates.Templates[i].ID
			templatesjson[i].Name = templates.Templates[i].Name
			templatesjson[i].Logo = xonglURL + "/" + logoRegex.FindStringSubmatch(templates.Templates[i].Template.String())[1]
		}
		return c.JSON(http.StatusOK, &templatesGet{
			Templates: templatesjson,
		})
	}
	return c.JSON(http.StatusBadRequest, errorMsg(err))
}

func getTemplate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		one := c.Get("one").(*goca.Controller)
		templateinfo, err := one.Template(id).Info(true, false)
		if err == nil {
			logoPath := logoRegex.FindStringSubmatch(templateinfo.Template.Template.String())[1]

			vcpuMin := vcpuRegex.FindStringSubmatch(templateinfo.Template.Template.String())[1]
			vcpuMax := vcpuRegex.FindStringSubmatch(templateinfo.Template.Template.String())[2]

			memoryMin := memoryRegex.FindStringSubmatch(templateinfo.Template.Template.String())[1]
			memoryMax := memoryRegex.FindStringSubmatch(templateinfo.Template.Template.String())[2]

			diskSize := disksizeRegex.FindStringSubmatch(templateinfo.Template.Template.String())[1]

			minVCPU, _ := strconv.Atoi(vcpuMin)
			maxVCPU, _ := strconv.Atoi(vcpuMax)

			minMEMORY, _ := strconv.Atoi(memoryMin)
			maxMEMORY, _ := strconv.Atoi(memoryMax)

			diskMin, _ := strconv.Atoi(diskSize)

			templatejson := templateDetGet{
				ID:        templateinfo.ID,
				Name:      templateinfo.Name,
				Logo:      xonglURL + "/" + logoPath,
				MinVCPU:   minVCPU,
				MaxVCPU:   maxVCPU,
				MinMEMORY: minMEMORY,
				MaxMEMORY: maxMEMORY,
				MinDisk:   diskMin,
				MaxDisk:   diskMax,
			}
			return c.JSON(http.StatusOK, templatejson)
		}
		return c.JSON(http.StatusBadRequest, errorMsg(err))
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid template ID[tid should be an integer]"})
}
