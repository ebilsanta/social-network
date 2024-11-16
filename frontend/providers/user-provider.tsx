'use client';

import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react';
import { Session } from 'next-auth';
import { useSession } from 'next-auth/react';
import { UserAPI } from '@/lib/user-api';
import { User } from '@/types/user';

interface UserContextType {
  user: User | null;
  setUser: (user: User | null) => void;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

interface UserProviderProps {
  children: ReactNode;
}

export const UserProvider: React.FC<UserProviderProps> = ({ children }) => {
  const { data: session, status } = useSession();
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    if (status === 'authenticated' && session?.user) {
      fetchUserData(session);
    } else {
      setUser(null);
    }
  }, [session, status]);

  const fetchUserData = async (userSession: Session) => {
    try {
      const userData = await UserAPI.getUser(userSession.user.id as string);
      setUser(userData.data);
    } catch (getUserErr) {
      if (getUserErr instanceof Error && getUserErr.message === 'User not found') {
        const { id, email, name, image } = userSession.user;
        try {
          const newUser = await UserAPI.createUser({
            id: id as string,
            email: email as string,
            image: image as string,
            name: name as string,
            username: name as string,
          });
          setUser(newUser.data);
        } catch (createUserErr) {
          if (createUserErr instanceof Error) {
            console.error(createUserErr.message);
            setUser(null);
          }
        }
      } else {
        console.error(getUserErr);
        setUser(null);
      }
    }
  };

  return <UserContext.Provider value={{ user, setUser }}>{children}</UserContext.Provider>;
};

export const useUser = (): UserContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
};
