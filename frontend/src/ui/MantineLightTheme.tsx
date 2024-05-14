import { MantineTheme, createTheme } from "@mantine/core";


export const MantineLightTheme = createTheme({
    scale: 1,
    fontSmoothing: true,
    focusRing: "auto",
    // colorScheme: 'light',
    white: "#fff",
    black: "#000",
    colors: {"dark": [
      "#C9C9C9",
      "#b8b8b8",
      "#828282",
      "#696969",
      "#424242",
      "#3b3b3b",
      "#2e2e2e",
      "#242424",
      "#1f1f1f",
      "#141414"
    ],
    "gray": [
      "#f8f9fa",
      "#f1f3f5",
      "#e9ecef",
      "#dee2e6",
      "#ced4da",
      "#adb5bd",
      "#868e96",
      "#495057",
      "#343a40",
      "#212529"
    ]
  },
  primaryShade: {
    light: 6,
    dark: 8
  },
    fontFamily: "Redhat, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji",
    // lineHeight: CSSProperties['lineHeight'];
    // transitionTimingFunction: CSSProperties['transitionTimingFunction'];
    // fontFamilyMonospace: CSSProperties['fontFamily'];
    // primaryColor: string;
  
    // fontSizes: Record<'xs' | 'sm' | 'md' | 'lg' | 'xl', number>;
    // spacing: Record<'xs' | 'sm' | 'md' | 'lg' | 'xl', number>;
    "respectReducedMotion": true,
    "headings": {
      "fontFamily": "Redhat, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji",
      "fontWeight": "700",
      "textWrap": "wrap",
      "sizes": {
        "h1": {
          "fontSize": "calc(2.125rem * var(--mantine-scale))",
          "lineHeight": "1.3"
        },
        "h2": {
          "fontSize": "calc(1.625rem * var(--mantine-scale))",
          "lineHeight": "1.35"
        },
        "h3": {
          "fontSize": "calc(1.375rem * var(--mantine-scale))",
          "lineHeight": "1.4"
        },
        "h4": {
          "fontSize": "calc(1.125rem * var(--mantine-scale))",
          "lineHeight": "1.45"
        },
        "h5": {
          "fontSize": "calc(1rem * var(--mantine-scale))",
          "lineHeight": "1.5"
        },
        "h6": {
          "fontSize": "calc(0.875rem * var(--mantine-scale))",
          "lineHeight": "1.5"
        }
      }
  },
  "shadows": {
    "xs": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), 0 calc(0.0625rem * var(--mantine-scale)) calc(0.125rem * var(--mantine-scale)) rgba(0, 0, 0, 0.1)",
    "sm": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(0.625rem * var(--mantine-scale)) calc(0.9375rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.4375rem * var(--mantine-scale)) calc(0.4375rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale))",
    "md": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(1.25rem * var(--mantine-scale)) calc(1.5625rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.625rem * var(--mantine-scale)) calc(0.625rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale))",
    "lg": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(1.75rem * var(--mantine-scale)) calc(1.4375rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.75rem * var(--mantine-scale)) calc(0.75rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale))",
    "xl": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(2.25rem * var(--mantine-scale)) calc(1.75rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(1.0625rem * var(--mantine-scale)) calc(1.0625rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale))"
  },
  "radius": {
    "xs": "calc(0.125rem * var(--mantine-scale))",
    "sm": "calc(0.25rem * var(--mantine-scale))",
    "md": "calc(0.5rem * var(--mantine-scale))",
    "lg": "calc(1rem * var(--mantine-scale))",
    "xl": "calc(2rem * var(--mantine-scale))"
  },
  
})