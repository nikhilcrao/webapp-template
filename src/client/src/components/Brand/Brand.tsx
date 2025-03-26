import { IconWallet } from '@tabler/icons-react';
import { Anchor, Group } from '@mantine/core';

export function Brand({ brandName: brand }: { brandName: string }) {
    return (
        <Anchor href="/">
            <Group>
                <IconWallet />
                <span style={{ fontVariant: 'small-caps' }}>
                    <strong>{brand}</strong>
                </span>
            </Group >
        </Anchor>
    );
}