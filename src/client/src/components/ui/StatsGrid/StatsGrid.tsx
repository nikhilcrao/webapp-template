import { Group, Paper, SimpleGrid, Text } from '@mantine/core';
import classes from './StatsGrid.module.css';

const data = [
  { title: 'Income', value: '$2,303', transactions: 23, color: 'teal' },
  { title: 'Expenses', value: '$4,145', transactions: 53, color: 'red' },
];

export function StatsGrid() {
  const stats = data.map((stat) => {

    return (
      <Paper withBorder p="sm" radius="md" key={stat.title}>
        <Group grow wrap="nowrap" justify="apart">
          <div>
            <Text c="dimmed" tt="uppercase" fw={500} fz="md" className={classes.label}>
              {stat.title}
            </Text>
            <Text fw={700} fz="xl" c={stat.color}>
              {stat.value}
            </Text>
          </div>
        </Group>
        <Text c="dimmed" fz="sm">
          {stat.transactions} {'transactions'}
        </Text>
      </Paper>
    );
  });

  return (
    <div className={classes.root}>
      <SimpleGrid cols={2}> {stats}</SimpleGrid>
    </div >
  );
}