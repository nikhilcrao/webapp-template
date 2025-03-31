import {
  Anchor,
  Button,
  Checkbox,
  Divider,
  Group,
  Paper,
  PaperProps,
  PasswordInput,
  Stack,
  Text,
  TextInput,
} from '@mantine/core';
import { useForm } from "@mantine/form";
import { upperFirst, useToggle } from "@mantine/hooks";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../contexts/AuthContext";
import { useGoogleLogin } from '@react-oauth/google';
import { GoogleButton } from "./GoogleButton";


export function AuthenticationForm(props: PaperProps) {
  const navigate = useNavigate();
  const authState = useAuth();

  const [type, toggle] = useToggle(['Login', 'Register']);
  const form = useForm({
    initialValues: {
      email: '',
      name: '',
      password: '',
      confirm_password: '',
      terms: true,
    },
    validate: {
      email: (val) => (/^\S+@\S+$/.test(val) ? null : 'invalid email'),
      password: (val) => (val.length <= 6 ? 'password should include at least 6 characters' : null),
    },
  });

  /*
  const queryParams = new URLSearchParams(location.search);
  const code = queryParams.get("code");
  */

  if (authState?.isAuthenticated && !authState?.isLoading) {
    navigate("/");
  }

  /*
  useEffect(() => {
    if (code) {
      const processCallback = async () => {
        try {
          const success = await authState?.processGoogleCallback(code);
          if (success) {
            navigate("/");
          }
        } catch (error) {
          console.error(error);
        }
      };

      processCallback();
    }
  }, [code, authState?.processGoogleCallback, navigate]);
  */


  const handleFormSubmit = async (data: any) => {
    try {
      if (type.toString() === 'Register') {
        if (data.password != data.confirm_password) {
          console.error("passwords do not match");
          return;
        }
        const success = await authState?.registerUser(data.name, data.email, data.password, data.confirm_password);
        if (success) {
          navigate("/");
        }
      } else {
        const success = await authState?.loginWithEmail(data.email, data.password);
        if (success) {
          navigate("/");
        }
      }
    } catch (error) {
      console.error(error);
    }
  };

  /*
  const handleGoogleLoginOld = async () => {
    try {
      await authState?.loginWithGoogle();
    } catch (error) {
      console.error(error);
    }
  }
    */

  const googleLogin = useGoogleLogin({
    onSuccess: codeResponse => {
      authState?.processGoogleCallback(codeResponse.code);
    },
    flow: 'auth-code',
    ux_mode: 'popup',
  });

  return (
    <Paper radius="md" p="xl" withBorder {...props}>
      <Text size="lg" fw={500}>
        {type}
      </Text>

      <Group grow mb="md" mt="md">
        <GoogleButton radius="xl" onClick={() => googleLogin()}>Google</GoogleButton>
      </Group>

      <Divider label="Or continue with email" labelPosition="center" my="lg" />

      <form onSubmit={form.onSubmit(handleFormSubmit)}>
        <Stack>
          {type === 'Register' && (
            <TextInput
              label="Name"
              value={form.values.name}
              onChange={(event) => form.setFieldValue('name', event.currentTarget.value)}
              radius="md"
            />
          )}

          <TextInput
            required
            label="Email"
            value={form.values.email}
            onChange={(event) => form.setFieldValue('email', event.currentTarget.value)}
            error={form.errors.email && 'Invalid email'}
            radius="md"
          />

          <PasswordInput
            required
            label="Password"
            value={form.values.password}
            onChange={(event) => form.setFieldValue('password', event.currentTarget.value)}
            error={form.errors.password && 'Password should include at least 6 characters'}
            radius="md"
          />

          {type === 'Register' && (
            <PasswordInput
              required
              label="Confirm Password"
              value={form.values.confirm_password}
              onChange={(event) => form.setFieldValue('confirm_password', event.currentTarget.value)}
              error={form.errors.password && 'Passwords must match'}
              radius="md"
            />
          )}

          {type === 'Register' && (
            <Checkbox
              label="I accept terms and conditions"
              checked={form.values.terms}
              onChange={(event) => form.setFieldValue('terms', event.currentTarget.checked)}
            />
          )}
        </Stack>

        <Group justify="space-between" mt="xl">
          <Anchor component="button" type="button" c="dimmed" onClick={() => toggle()} size="xs">
            {type === 'Register'
              ? 'Already have an account? Login'
              : "Don't have an account? Register"}
          </Anchor>
          <Button type="submit" radius="xl">
            {upperFirst(type)}
          </Button>
        </Group>
      </form>
    </Paper >
  );
}