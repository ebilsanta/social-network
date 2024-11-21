import React from 'react';
import { redirect } from 'next/navigation';
import { IconSearch } from '@tabler/icons-react';
import {
  Autocomplete,
  AutocompleteProps,
  Avatar,
  Group,
  Popover,
  Text,
  Tooltip,
  UnstyledButton,
} from '@mantine/core';
import { useUserSearch } from '@/app/_components/Navbar/UserSearch/useUserSearch';
import classes from '../Navbar.module.css';

interface UserSearchProps {
  isSmallScreen: boolean | undefined;
}

export const UserSearch = ({ isSmallScreen }: UserSearchProps) => {
  const { usersData, query, setQuery } = useUserSearch();
  const renderAutocompleteOption: AutocompleteProps['renderOption'] = ({ option }) => (
    <UnstyledButton onClick={() => redirect(`/${option.value}`)} style={{ width: '100%' }}>
      <Group gap="sm">
        <Avatar src={usersData[option.value].image} size={36} radius="xl" />
        <div>
          <Text size="sm">{option.value}</Text>
          <Text size="xs" opacity={0.6}>
            {usersData[option.value].name}
          </Text>
        </div>
      </Group>
    </UnstyledButton>
  );
  const autocompleteProps = {
    value: query,
    onChange: setQuery,
    data: Object.keys(usersData).map((username) => ({ value: username })),
    renderOption: renderAutocompleteOption,
    maxDropdownHeight: 400,
    placeholder: 'Search for users',
  };
  if (isSmallScreen) {
    return (
      <Popover width={300} position="right" withArrow shadow="md">
        <Popover.Target>
          <Tooltip label="Search" position="right" transitionProps={{ duration: 0 }}>
            <UnstyledButton className={classes.link} data-minimized="true">
              <IconSearch
                className={classes.linkIcon}
                width={25}
                height={25}
                stroke={1.5}
                data-minimized="true"
              />
            </UnstyledButton>
          </Tooltip>
        </Popover.Target>
        <Popover.Dropdown>
          <Autocomplete {...autocompleteProps} comboboxProps={{ withinPortal: false }} />
        </Popover.Dropdown>
      </Popover>
    );
  }
  return <Autocomplete {...autocompleteProps} />;
};
