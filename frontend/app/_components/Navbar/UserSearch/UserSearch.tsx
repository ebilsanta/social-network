import { Autocomplete, AutocompleteProps, Avatar, Group, Text } from '@mantine/core';
import { useUserSearch } from '@/app/_components/Navbar/UserSearch/useUserSearch';

export const UserSearch = () => {
  const { usersData, query, setQuery } = useUserSearch();
  const renderAutocompleteOption: AutocompleteProps['renderOption'] = ({ option }) => (
    <Group gap="sm">
      <Avatar src={usersData[option.value].image} size={36} radius="xl" />
      <div>
        <Text size="sm">{option.value}</Text>
        <Text size="xs" opacity={0.6}>
          {usersData[option.value].name}
        </Text>
      </div>
    </Group>
  );
  return (
    <Autocomplete
      value={query}
      onChange={setQuery}
      data={Object.keys(usersData).map((username) => ({ value: username }))}
      renderOption={renderAutocompleteOption}
      maxDropdownHeight={400}
      placeholder="Search for users"
    />
  );
};
