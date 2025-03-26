import { IconWallet } from '@tabler/icons-react';
import { Group } from '@mantine/core';

export function Brand({ brandName: brand }: { brandName: string }) {
    return (
        <Group>
            <IconWallet />
            <span style={{ fontVariant: 'small-caps' }}>
                <strong>{brand}</strong>
            </span>
        </Group >

    );
}