import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/charts/styles.css';

import { AppShell, Burger, Stack, Loader, MantineProvider, Flex } from '@mantine/core';
import { HeaderMenu } from './components/HeaderMenu/HeaderMenu';


export default function App() {
  return (
    <MantineProvider defaultColorScheme="auto">
      <AppShell
        header={{ height: 60 }}
        navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
        padding="md"
        transitionDuration={500}
        transitionTimingFunction="ease"
      >
        <AppShell.Header>
          <HeaderMenu />
        </AppShell.Header>

        <AppShell.Navbar p="md">Navbar</AppShell.Navbar>

        <AppShell.Main>
          <Stack>
            <div>Main</div>
            <Loader color="blue" />
          </Stack>
        </AppShell.Main>
      </AppShell>
    </MantineProvider>
  );
}