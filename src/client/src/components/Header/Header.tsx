import { Burger, Container } from '@mantine/core';
import { Brand } from '../Brand/Brand';
import classes from './Header.module.css';

export interface HeaderProps {
  opened: boolean;
  toggle?: () => void;
}

export function Header(props: HeaderProps) {
  return (
    <header className={classes.header}>
      <Container fluid>
        <div className={classes.inner}>
          <Brand brandName='Webapp Template' />
          {props.toggle ? <Burger opened={props.opened} onClick={props.toggle} size="sm" hiddenFrom="md" /> : null}
        </div>
      </Container>
    </header>
  );
}