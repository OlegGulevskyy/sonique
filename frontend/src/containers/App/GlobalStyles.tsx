import { Global } from "@mantine/core";

export const GlobalStyles = () => {
  return (
    <Global
      styles={(theme) => ({
        body: {
          ...theme.fn.fontStyles(),
          backgroundColor: "transparent",
          color:
            theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,
          lineHeight: theme.lineHeight,
          border: "1px solid red",
        },
      })}
    />
  );
};
