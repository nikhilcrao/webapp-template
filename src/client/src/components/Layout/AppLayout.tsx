import { AppShell } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Outlet } from 'react-router-dom';
import { Header } from '../Header/Header';
import { Navbar } from '../Navbar/Navbar';

export function AppLayout() {
  const [opened, { toggle }] = useDisclosure();

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{
        width: 300,
        breakpoint: "md",
        collapsed: { mobile: !opened },
      }}
      padding="md"
    >

      <AppShell.Header>
        <Header opened={opened} toggle={toggle} />
      </AppShell.Header>

      <AppShell.Navbar>
        <Navbar />
      </AppShell.Navbar>

      <AppShell.Main>
        <Outlet />
      </AppShell.Main>

    </AppShell>
  );
}