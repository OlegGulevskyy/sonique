import { Button, MantineProvider } from "@mantine/core";

function App() {
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <div>
        <Button>Click me</Button>
      </div>
    </MantineProvider>
  );
}

export default App;
