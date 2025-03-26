import { Burger, Center, Container, Group, Menu } from '@mantine/core';
import { Brand } from '../Brand/Brand';
import classes from './Header.module.css';


export function Header({ opened, toggle }: { opened: boolean, toggle: () => void }) {
  return (
    <header className={classes.header}>
      <Container size="md">
        <div className={classes.inner}>
          <Brand brandName='Webapp Template' />
          <Burger opened={opened} onClick={toggle} size="sm" hiddenFrom="sm" />
        </div>
      </Container>
    </header>
  );
}