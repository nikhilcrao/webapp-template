// Import styles of packages that you've installed.
// All packages except `@mantine/hooks` require styles imports
import '@mantine/core/styles.css';

import { MantineProvider, AppShell, Burger } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';

export default function App() {
  const [opened, { toggle }] = useDisclosure();

  return (
    <MantineProvider>
      <AppShell
        header={{ height: 60 }}
        navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
        padding="md"
        transitionDuration={500}
        transitionTimingFunction="ease"
      >
        <AppShell.Header>
          <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm"></Burger>
          <div>Logo</div>
        </AppShell.Header>

        <AppShell.Navbar p="md">Navbar</AppShell.Navbar>

        <AppShell.Main>Main</AppShell.Main>
      </AppShell>
    </MantineProvider>
  );
}