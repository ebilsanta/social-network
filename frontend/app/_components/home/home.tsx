'use client';

import { Feed } from '@/app/_components/Home/Feed/Feed';
import { Navbar } from '@/app/_components/Home/Navbar/Navbar';
import classes from './Home.module.css';

export function Home() {
  return (
    <div className={classes.container}>
      <Navbar /> <Feed />
    </div>
  );
}
