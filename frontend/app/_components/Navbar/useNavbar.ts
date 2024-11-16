'use client';

import { useEffect, useState } from 'react';
import { usePathname } from 'next/navigation';
import { IconCirclePlus, IconHome, IconUserCircle } from '@tabler/icons-react';
import { User } from '@/types/user';

export const useNavbar = (user: User) => {
  const [path, setPath] = useState('');
  const pathname = usePathname();
  useEffect(() => {
    setPath(pathname);
  }, [pathname]);

  const data = [
    { link: '/', label: 'Home', icon: IconHome },
    { link: '', label: 'Create', icon: IconCirclePlus },
    {
      link: user?.username ? `/${user.username}` : '',
      label: 'Profile',
      icon: IconUserCircle,
    },
  ];
  return {
    data,
    path,
  };
};
