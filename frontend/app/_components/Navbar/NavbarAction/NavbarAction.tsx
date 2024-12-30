import React from 'react';
import { IconHome2 } from '@tabler/icons-react';
import { Tooltip } from '@mantine/core';
import classes from '../Navbar.module.css';

interface NavbarLinkProps {
  icon: typeof IconHome2;
  label: string;
  onClick: ((event: React.MouseEvent) => void) | (() => void);
  active?: boolean;
  isSmallScreen?: boolean;
}

export const NavbarAction = ({ icon, label, onClick, active, isSmallScreen }: NavbarLinkProps) => {
  const dataMinimized = isSmallScreen || undefined;
  const dataActive = active || undefined;
  return (
    <Tooltip
      label={label}
      position="right"
      transitionProps={{ duration: 0 }}
      disabled={!isSmallScreen}
    >
      <a
        href="#"
        className={classes.link}
        data-minimized={dataMinimized}
        data-active={dataActive}
        onClick={onClick}
      >
        {React.createElement(icon as any, {
          className: classes.linkIcon,
          width: 25,
          height: 25,
          stroke: 1.5,
          'data-minimized': dataMinimized,
        })}
        {!isSmallScreen ? <span>{label}</span> : null}
      </a>
    </Tooltip>
  );
};
