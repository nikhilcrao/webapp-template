import {
    Text,
    Card,
    Paper,
    SimpleGrid,
    Badge,
    Group,
    Table,
    Drawer,
    Modal,
    ActionIcon,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconAutomation } from "@tabler/icons-react";
import { createContext, useState } from "react";

function formatIndianRupees(number: number) {
    return new Intl.NumberFormat('en-IN', {
        style: 'currency',
        currency: 'INR',
    }).format(number);
}

function StatsGrid({ income, expense }: { income: number, expense: number }) {
    return (
        <Paper shadow={"xs"} key={"statsGrid"}>
            <SimpleGrid cols={2} spacing="md">
                <Card>
                    <Text fz={"xs"} c={"dimmed"} tt={"uppercase"}>Income</Text>
                    <Text fz={"lg"} fw={500} c="green">{formatIndianRupees(income)}</Text>
                </Card>
                <Card>
                    <Text fz={"xs"} c={"dimmed"} tt={"uppercase"}>Expenses</Text>
                    <Text fz={"lg"} fw={500} c="red">{formatIndianRupees(expense)}</Text>
                </Card>
            </SimpleGrid>
        </Paper>
    );
}

function MenuDrawer({ opened, close, text }: { opened: boolean, close: () => void, text: any }) {
    return (
        <Drawer
            opened={opened}
            onClose={close}
            key={"drawer"}
            title="Menu"
            position="bottom">
            <Text>{text}</Text>
        </Drawer>
    );
}

interface DrawerContextType {
    drawerData: any;
    setDrawerData: React.Dispatch<React.SetStateAction<any>>;
}

const DrawerContext = createContext<DrawerContextType | null>(null);

const data = [
    {
        id: "abc",
        description: "Zomato",
        amount: 4203,
        category: "Dining",
    },
    {
        id: "234",
        description: "Indigo",
        amount: 12223,
        category: "Travel",
        merchant: "Indigo",
    },
    {
        id: "345",
        description: "UPI Payment",
        amount: 283,
    },
];

export function Dashboard() {
    const [drawerData, setDrawerData] = useState<string | null>(null);
    const [opened, { open, close }] = useDisclosure(false);
    const [modalOpened, { open: openModal, close: closeModal }] = useDisclosure(false);

    const rows = data.map((row) => (
        <Table.Tr key={row.id}>
            <Table.Td
                width={"90%"}
                onClick={() => {
                    setDrawerData(row.id);
                    open();
                }}
            >
                <Text fz={"md"} fw={300} truncate="end">{row.description}</Text>
                <Group>
                    <Text fz={"xs"} c={"dimmed"}>{formatIndianRupees(row.amount)}</Text>

                    {row.category &&
                        <Badge size={"xs"} autoContrast>{row.category}</Badge>
                    }
                    {row.merchant &&
                        <Badge size={"xs"} autoContrast>{row.merchant}</Badge>
                    }
                </Group>
            </Table.Td>
            <Table.Td>
                <ActionIcon
                    variant="subtle"
                    aria-label="Create Rule"
                    onClick={() => {
                        setDrawerData(row.id);
                        openModal();
                    }}
                >
                    <IconAutomation stroke={1.5} />
                </ActionIcon>
            </Table.Td>
        </Table.Tr >
    ));

    return (
        <DrawerContext.Provider value={{ drawerData, setDrawerData }}>
            <Modal opened={modalOpened} onClose={closeModal} key={"modalKey"} centered>
                <Text>Modal: {drawerData}</Text>
            </Modal>
            <MenuDrawer opened={opened} close={close} text={drawerData} />
            <StatsGrid income={232398} expense={213423} />
            <Table p="md" mt="md" key={"txnTable"}>
                <Table.Tbody>{rows}{rows}{rows}{rows}{rows}</Table.Tbody>
            </Table>
        </DrawerContext.Provider>
    );
}