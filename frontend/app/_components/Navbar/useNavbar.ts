'use client';

import { useEffect, useState } from 'react';
import { redirect, usePathname } from 'next/navigation';
import {
  IconCirclePlus,
  IconHome,
  IconLogout,
  IconSwitchHorizontal,
  IconUserCircle,
} from '@tabler/icons-react';
import { signIn, signOut } from 'next-auth/react';
import { User } from '@/types/user';

export const useNavbar = (user: User) => {
  const [path, setPath] = useState('');
  const pathname = usePathname();
  useEffect(() => {
    setPath(pathname);
  }, [pathname]);

  const handleLinkClick = (event: React.MouseEvent, link: string) => {
    event.preventDefault();
    if (link) {
      redirect(link);
    }
  };

  const navbarLinks = [
    {
      active: path === '/',
      label: 'Home',
      icon: IconHome,
      onClick: (event: React.MouseEvent) => handleLinkClick(event, '/'),
    },
    {
      active: path === '/create',
      label: 'Create',
      icon: IconCirclePlus,
      onClick: (event: React.MouseEvent) => handleLinkClick(event, '/create'),
    },
    {
      active: path === `/${user.username}`,
      label: 'Profile',
      icon: IconUserCircle,
      onClick: (event: React.MouseEvent) => handleLinkClick(event, `/${user.username}`),
    },
  ];

  const navbarActions = [
    {
      active: false,
      label: 'Change account',
      icon: IconSwitchHorizontal,
      onClick: () => signIn(),
    },
    {
      active: false,
      label: 'Logout',
      icon: IconLogout,
      onClick: () => signOut(),
    },
  ];

  return {
    navbarLinks,
    navbarActions,
  };
};
