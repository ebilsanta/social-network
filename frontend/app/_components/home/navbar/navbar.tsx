import {
  IconCirclePlus,
  IconHome,
  IconLogout,
  IconSearch,
  IconSwitchHorizontal,
  IconUserCircle,
} from '@tabler/icons-react';
import { signOut } from 'next-auth/react';
import { Code, Group, rem, TextInput } from '@mantine/core';
import { MantineLogo } from '@mantinex/mantine-logo';
import { UserButton } from '@/app/_components/home/navbar/user-button/user-button';
import { useUser } from '@/providers/user-provider';
import classes from './navbar.module.css';

export function Navbar() {
  const { user } = useUser();
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
          <MantineLogo size={28} />
          <Code fw={700}>v3.1.2</Code>
        </Group>
        <TextInput
          placeholder="Search for users"
          size="xs"
          leftSection={<IconSearch style={{ width: rem(12), height: rem(12) }} stroke={1.5} />}
          styles={{ section: { pointerEvents: 'none' } }}
          mb="sm"
        />
        {links}
      </div>

      <div className={classes.footer}>
        <div className={classes.section}>
          <UserButton user={user} />
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
}
