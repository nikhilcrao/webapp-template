import { AppShell } from '@mantine/core';
import { AuthenticationForm } from "../../components/auth/AuthForm/AuthForm";
import { Header } from '../../components/layout/Header/Header';

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
