import { AppShell } from '@mantine/core';
import { AuthenticationForm } from "../Auth/AuthenticationForm";
import { Header } from '../Header/Header';

export function LoginPage() {
    return (
        <AppShell header={{ height: 60 }} padding="md">

            <AppShell.Header>
                <Header opened={false} />
            </AppShell.Header>

            <AppShell.Main>
                <AuthenticationForm />
            </AppShell.Main>
        </AppShell>
    );
}
