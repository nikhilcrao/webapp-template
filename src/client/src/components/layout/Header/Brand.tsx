import { IconWallet } from '@tabler/icons-react';
import { Anchor, Group } from '@mantine/core';

export function Brand() {
    return (
        <Anchor href="/">
            <Group>
                <IconWallet />
                <span style={{ fontVariant: 'small-caps' }}>
                    <strong>Budget Buddy</strong>
                </span>
            </Group >
        </Anchor>
    );
}