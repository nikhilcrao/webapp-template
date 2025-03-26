import { useState } from 'react';
import { IconChevronRight } from '@tabler/icons-react';
import { Box, Collapse, Group, ThemeIcon, UnstyledButton } from '@mantine/core';
import { Link } from 'react-router-dom';
import classes from './LinksGroup.module.css';

export interface LinksGroupProps {
  icon: React.FC<any>;
  label: string;
  initiallyOpened?: boolean;
  link?: string;
  linksGroup?: { label: string; link: string }[];
}

export function LinksGroup({ icon: Icon, label, initiallyOpened, link, linksGroup }: LinksGroupProps) {
  const [opened, setOpened] = useState(initiallyOpened || false);

  const hasLink = (link === null);
  const hasLinkGroup = Array.isArray(linksGroup);
  const items = (hasLinkGroup ? linksGroup : []).map((link) => (
    <Link to={link.link} className={classes.link} key={link.link}>
      {link.label}
    </Link>
  ));

  return (
    <>
      <UnstyledButton onClick={() => setOpened((opened: boolean) => !opened)} className={classes.control}>
        <Group justify="space-between" gap={0}>
          <Box style={{ display: 'flex', alignItems: 'center' }}>
            <ThemeIcon variant="light" size={30}>
              <Icon size={18} />
            </ThemeIcon>
            {hasLink ? (
              <Link to={link}>
                <Box ml="md">{label}</Box>
              </Link>
            ) : (
              <Box ml="md">{label}</Box>
            )}
          </Box>

          {hasLinkGroup && (
            <IconChevronRight
              className={classes.chevron}
              stroke={1.5}
              size={16}
              style={{ transform: opened ? 'rotate(-90deg)' : 'none' }}
            />
          )}
        </Group>
      </UnstyledButton>

      {hasLinkGroup ? <Collapse in={opened}>{items}</Collapse> : null}
    </>
  );
}