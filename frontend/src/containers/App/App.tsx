import { MainInput } from "@/components/MainInput";
import { MantineProvider } from "@mantine/core";
import { GlobalStyles } from "./GlobalStyles";

export const App = () => {
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <GlobalStyles />
      <MainInput />
    </MantineProvider>
  );
};
