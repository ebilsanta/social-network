import {
  IconCirclePlus,
  IconHome,
  IconLogout,
  IconSwitchHorizontal,
  IconUserCircle,
} from '@tabler/icons-react';
import { signOut } from 'next-auth/react';
import { Code, Group, Space } from '@mantine/core';
import { MantineLogo } from '@mantinex/mantine-logo';
import { ProfileCard } from '@/app/_components/Navbar/ProfileCard/ProfileCard';
import { UserSearch } from '@/app/_components/Navbar/UserSearch/UserSearch';
import { User } from '@/types/user';
import classes from './Navbar.module.css';

interface NavbarProps {
  user: User | null;
}

export const Navbar = ({ user }: NavbarProps) => {
  const data = [
    { link: '', label: 'Home', icon: IconHome },
    { link: '', label: 'Create', icon: IconCirclePlus },
    {
      link: '',
      label: 'Profile',
      icon: IconUserCircle,
    },
  ];

  const links = data.map((item) => (
    <a
      className={classes.link}
      href={item.link}
      key={item.label}
      onClick={(event) => {
        event.preventDefault();
      }}
    >
      <item.icon className={classes.linkIcon} stroke={1.5} />
      <span>{item.label}</span>
    </a>
  ));

  return (
    <nav className={classes.navbar}>
      <div className={classes.navbarMain}>
        <Group className={classes.header} justify="space-between">
          <MantineLogo size={36} />
          <Code fw={700}>v3.1.2</Code>
        </Group>
        <UserSearch />
        <Space h="sm" />
        {links}
      </div>

      <div className={classes.footer}>
        <div className={classes.section}>
          <ProfileCard user={user} />
        </div>
        <a href="#" className={classes.link} onClick={(event) => event.preventDefault()}>
          <IconSwitchHorizontal className={classes.linkIcon} stroke={1.5} />
          <span>Change account</span>
        </a>

        <a href="#" className={classes.link} onClick={() => signOut()}>
          <IconLogout className={classes.linkIcon} stroke={1.5} />
          <span>Logout</span>
        </a>
      </div>
    </nav>
  );
};
