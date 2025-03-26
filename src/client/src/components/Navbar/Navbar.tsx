import {
  IconAdjustments,
  IconGauge,
  IconNotes,
  IconPresentationAnalytics,
} from '@tabler/icons-react';

import { IconLogout } from '@tabler/icons-react';
import { ScrollArea } from '@mantine/core';
import { LinksGroup, LinksGroupProps } from './LinksGroup';
import classes from './Navbar.module.css';

export const navItems = {
  items: [
    {
      label: "Dashboard",
      icon: IconGauge,
      link: "/",
    },
    {
      label: "Forecasts",
      icon: IconNotes,
      link: "/forecasts",
      linksGroup: [
        { label: "Overview", link: "/forecasts/overview" },
        { label: "Forecasts", link: "/forecasts/forecasts" },
        { label: "Outlook", link: "/forecasts/outlook" },
        { label: "Real time", link: "/forecasts/real-time" },
      ],
    },
    {
      label: "Analytics",
      icon: IconPresentationAnalytics,
      link: "/analytics",
      linksGroup: [
        { label: "Reports", link: "/analytics/reports" },
      ],
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
        {/* TODO: handle logout */}
        <a href="/logout" className={classes.link} onClick={(event) => event.preventDefault()}>
          <IconLogout className={classes.linkIcon} stroke={1.5} />
          <span>Logout</span>
        </a>
      </div>
    </nav>
  );
}