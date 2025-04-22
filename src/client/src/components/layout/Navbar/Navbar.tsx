import {
  IconAdjustments,
  IconAutomation,
  IconCategory,
  IconDashboard,
  IconFileImport,
  IconShoppingBag,
} from '@tabler/icons-react';

import { IconLogout } from '@tabler/icons-react';
import { ScrollArea, Anchor } from '@mantine/core';
import { LinksGroup } from './LinksGroup';
import classes from './Navbar.module.css';

/*
{
      label: "Transactions",
      icon: IconNotes,
      link: "/forecasts",
      linksGroup: [
        { label: "Overview", link: "/forecasts/overview" },
        { label: "Forecasts", link: "/forecasts/forecasts" },
        { label: "Outlook", link: "/forecasts/outlook" },
        { label: "Real time", link: "/forecasts/real-time" },
      ],
    },
*/
export const navItems = {
  items: [
    {
      label: "Dashboard",
      icon: IconDashboard,
      link: "/",
    },
    {
      label: "Import",
      icon: IconFileImport,
      link: "/transactions/import",
    },
    {
      label: "Automation",
      icon: IconAutomation,
      link: "/rules",
    },
    {
      label: "Categories",
      icon: IconCategory,
      link: "/categories",
    },
    {
      label: "Merchants",
      icon: IconShoppingBag,
      link: "/merchants",
    },
    {
      label: "Settings",
      icon: IconAdjustments,
      link: "/settings",
    },
  ]
};

export function Navbar() {
  const links = navItems.items.map((item) => <LinksGroup {...item} key={item.label} />);

  return (
    <nav className={classes.navbar}>
      <ScrollArea className={classes.links}>
        <div className={classes.linksInner}>{links}</div>
      </ScrollArea>

      <div className={classes.footer}>
        <Anchor href="/logout" className={classes.link}>
          <IconLogout className={classes.linkIcon} stroke={1.5} />
          <span>Logout</span>
        </Anchor>
      </div>
    </nav>
  );
}