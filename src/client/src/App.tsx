import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/charts/styles.css';

import { Anchor, AppShell, MantineProvider, Flex, Breadcrumbs } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';

import { Header } from './components/Header/Header';
import { Navbar } from './components/Navbar/Navbar';
import { Footer } from './components/Footer/Footer';
import { TableSort } from './components/TableSort/TableSort';
import { StatsGrid } from './components/StatsGrid/StatsGrid';


export default function App() {
  const [opened, { toggle }] = useDisclosure();
  
  const items = [
    { title: 'Mantine', href: '#' },
    { title: 'Mantine hooks', href: '#' },
    { title: 'use-id', href: '#' },
  ].map((item, index) => (
    <Anchor href={item.href} key={index}>
      {item.title}
    </Anchor>
  ));

  return (
    <MantineProvider defaultColorScheme="auto">
      <AppShell
        header={{ height: 60 }}
        navbar={{
          width: 300,
          breakpoint: "sm",
          collapsed: { mobile: !opened }, 
        }}
        padding="md"
      >
        <AppShell.Header>
          <Header
            opened={opened}
            onMenuOpen={toggle}
          />
        </AppShell.Header>

        <AppShell.Navbar>
          <Navbar />
        </AppShell.Navbar>

        <AppShell.Main>
          <Flex direction="column">
            <Breadcrumbs>{ items }</Breadcrumbs>
            <StatsGrid />
            <TableSort />
          </Flex>
        </AppShell.Main>

        <AppShell.Footer>
          <Footer />
        </AppShell.Footer>

      </AppShell>
    </MantineProvider>
  );
}