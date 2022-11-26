import { Button, MantineProvider } from "@mantine/core";
import { GlobalStyles } from "./GlobalStyles";

export const App = () => {
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <GlobalStyles />
      <div>
        <Button>Click me</Button>
      </div>
    </MantineProvider>
  );
};
