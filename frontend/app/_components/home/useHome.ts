'use client';

import { useUser } from '@/providers/user-provider';

export const useHome = () => {
  const { user, setUser } = useUser();
  return { user, setUser };
};
