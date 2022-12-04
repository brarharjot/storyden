import { defineStyleConfig, extendTheme } from "@chakra-ui/react"

const Button = defineStyleConfig({
  variants: {
    outline: {
      border: "1px solid",
      borderColor: "primary.base",
      borderRadius: "12px",
      fontWeight: "normal",
      _focus: {
        outline: "2px solid",
        outlineColor: "primary.active",
        transition: "outline 0.3s",
      },
    },
    solid: {
      bg: "primary.base",
      color: "primary.contrast",
      borderRadius: "12px",
      fontWeight: "normal",
      _hover: {
        bg: "primary.hover",
      },
      _active: {
        bg: "primary.active",
      },
      _focus: {
        outline: "2px solid",
        outlineColor: "primary.active",
        transition: "outline 0.3s",
      },
    },
  },
  defaultProps: {
    size: "md",
    variant: "solid",
  },
})

export const theme = extendTheme({
  components: {
    Button,
  },
  colors: {
    primary: {
      base: "#044162",
      hover: "#03314A",
      active: "#02273B",
      contrast: "#FFFFFF",
    },
    background: {
      base: "#001322",
    },
  },
  fonts: {
    heading: "p22-mackinac-pro",
    body: "halyard-display",
  },
})
