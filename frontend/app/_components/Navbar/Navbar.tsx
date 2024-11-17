import { redirect } from 'next/navigation';
import { Code, Group, Space } from '@mantine/core';
import { MantineLogo } from '@mantinex/mantine-logo';
import { NavbarAction } from '@/app/_components/Navbar/NavbarAction/NavbarAction';
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
  const { navbarActions, navbarLinks } = useNavbar(user);

  const links = navbarLinks.map((item) => (
    <NavbarAction
      key={item.label}
      icon={item.icon}
      label={item.label}
      onClick={item.onClick}
      active={item.active}
      isSmallScreen={isSmallScreen}
    />
  ));

  const actions = navbarActions.map((action, index) => (
    <NavbarAction
      key={index}
      icon={action.icon}
      label={action.label}
      onClick={action.onClick}
      active={action.active}
      isSmallScreen={isSmallScreen}
    />
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
        <ProfileCard user={user} isSmallScreen={isSmallScreen} />
        {actions}
      </div>
    </nav>
  );
};
