import { redirect } from 'next/navigation';
import { IconLogout, IconSwitchHorizontal } from '@tabler/icons-react';
import { signIn, signOut } from 'next-auth/react';
import { Code, Group, Space } from '@mantine/core';
import { MantineLogo } from '@mantinex/mantine-logo';
import { ProfileCard } from '@/app/_components/Navbar/ProfileCard/ProfileCard';
import { useNavbar } from '@/app/_components/Navbar/useNavbar';
import { UserSearch } from '@/app/_components/Navbar/UserSearch/UserSearch';
import { User } from '@/types/user';
import classes from './Navbar.module.css';

interface NavbarProps {
  user: User;
  isSmallScreen: boolean | undefined;
}

export const Navbar = ({ user, isSmallScreen }: NavbarProps) => {
  const { data, path } = useNavbar(user);

  const links = data.map((item) => (
    <a
      className={classes.link}
      href={item.link}
      key={item.label}
      data-active={item.link === path || undefined}
      data-minimized={isSmallScreen || undefined}
      onClick={(event) => {
        event.preventDefault();
        if (item.link) {
          redirect(item.link);
        }
      }}
    >
      <item.icon className={classes.linkIcon} stroke={1.5} />
      <span>{item.label}</span>
    </a>
  ));

  return (
    <nav className={classes.navbar} data-minimized={isSmallScreen || undefined}>
      <div className={classes.navbarMain}>
        <Group
          className={classes.header}
          justify="space-between"
          onClick={() => redirect('/')}
          style={{ cursor: 'pointer' }}
        >
          <MantineLogo size={36} />
          <Code fw={700}>v3.1.2</Code>
        </Group>
        <UserSearch />
        <Space h="sm" />
        {links}
      </div>

      <div className={classes.footer}>
        <ProfileCard user={user} />
        <a href="#" className={classes.link} onClick={() => signIn()}>
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
