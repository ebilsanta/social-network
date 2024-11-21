'use client';

import { ReactNode } from 'react';
import { useMediaQuery } from '@mantine/hooks';
import { Navbar } from '@/app/_components/Navbar/Navbar';
import { useUser } from '@/providers/user-provider';
import classes from './Layout.module.css';

export const Layout = ({ children }: { children: ReactNode }) => {
  const { user } = useUser();
  const isSmallScreen = useMediaQuery('(max-width: 860px)');

  return (
    <div
      className={classes.container}
      data-navbar={!!user || undefined}
      data-minimized={(user && isSmallScreen) || undefined}
    >
      {user && <Navbar user={user} isSmallScreen={isSmallScreen} />}
      <div className={classes.content}>{children}</div>
    </div>
  );
};
