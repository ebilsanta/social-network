'use client';

import { useState } from 'react';
import useSWR from 'swr';
import { useDebouncedValue } from '@mantine/hooks';
import { UserAPI } from '@/lib/user-api';
import { User } from '@/types/user';

const PAGE = 1;
const LIMIT = 10;
const DEBOUNCE_DELAY = 1000;

const fetchUsers = async (query: string) => {
  if (!query) return [];
  const response = await UserAPI.getUsers(query, PAGE, LIMIT);
  return response.data;
};

export const useUserSearch = () => {
  const [query, setQuery] = useState('');
  const [debouncedQuery] = useDebouncedValue(query, DEBOUNCE_DELAY);

  const { data, error, isValidating } = useSWR(
    debouncedQuery ? `/api/users/${debouncedQuery}` : null,
    () => fetchUsers(query)
  );

  const loading = isValidating || !data;

  const usersData = data
    ? data.reduce((acc: Record<string, User>, user: User) => {
        acc[user.username] = user;
        return acc;
      }, {})
    : {};

  return {
    query,
    setQuery,
    loading,
    usersData,
    error,
  };
};
