/*=========================================================================================
  File Name: sidebarItems.js
  Description: Sidebar Items list. Add / Remove menu items from here.
  ----------------------------------------------------------------------------------------
  Item Name: Xongl user panel
  Author: Xongl cloud private limited
==========================================================================================*/


export default [
  {
    url: "/home",
    name: "Dashboard",
    slug: "home",
    icon: "HomeIcon",
  },
  // {
  //   header: "Instances",
  //   icon: "PackageIcon",
  //   items: [
  //     {
  //       url: null,
  //       name: "Project",
  //       icon: "PlusIcon",
  //       submenu: [
  //         {
  //           url: '/apps/user/user-list',
  //           name: "List",
  //           slug: "app-user-list",

  //         },
  //       ]
  //     },
  //   ]
  // },
  {
    header: "Manage",
    icon: "LayersIcon",
    items: [
      {
        url: null,
        name: "Deploy VM",
        tag: "new",
        tagColor: "primary",
        icon: "ListIcon",
        submenu: [
          {
            url: '/select-template',
            name: "OS/Apps",
            slug: "select-template-page",
          },
        ]
      },    
    ]
  },
  {
    header: "Accounts",
    icon: "Edit3Icon",
    items: [
      // {
      //   url:  '/vm-usage',
      //   name: "Usage",
      //   icon: "PieChartIcon",
      //   slug: "vm-usage-page"
      // },
      // {
      //   url: '/vm-showback',
      //   icon: "MonitorIcon",
      //   name: "Showback",
      //   slug: "vm-showback-page",
      // },
      {
        url: 'https://www.xongl.com/client/clientarea.php?action=invoices',
        icon: "FileTextIcon",
        name: "Billing & Invoice",
        slug: "external",
        target: "_blank"
      },
    ]
  },
  {
    header: "Support",
    icon: "HeadPhoneIcon",
    items: [
          {
            url: 'https://www.xongl.com/client/index.php?rp=/knowledgebase',
            name: "Documentation",
            icon: "BookOpenIcon",
            slug: "external",
            target: "_blank"
          },
          {
            url: 'https://www.xongl.com/client/supporttickets.php',
            name: "Raise Support",
            icon: "LifeBuoyIcon",
            slug: "external",
            target: "_blank"
          },
        ]
  },  
]

