import { createTheme } from "@mantine/core"


export const AnicLightTheme = createTheme({

  // colorScheme: 'dark',
  // variantColorResolver: true,
  /** White color */
  white: "FBFCFF",

  /** Black color */
  black: "070B0D",

  /** Object of colors, key is color name, value is an array of at least 10 strings (colors) */
  colors: {
    "brand": ['#F0F', '#E0E', '#D0D', '#C0C', '#B0B', '#A0A', '#909', '#808', '#707', '#606'],
    // "dark": ["#f2f5ff", "#d4ddf7", "#b8c6f0", "#9bafe8", "#110b40", "#0c0831", "#080521", "#040312", "#000003", "#000"],
    // "dark": ["#f7f9fa", "#d2d6dc", "#adb5be", "#8993a1", "#759fff", "#577fd9", "#1e4090", "#121212", "#01216b", "#000"],,
    'bright-pink': ['#F0BBDD', '#ED9BCF', '#EC7CC3', '#ED5DB8', '#F13EAF', '#F71FA7', '#FF00A1', '#E00890', '#C50E82', '#AD1374'],
    "light": [
      "#fff1e8",
      "#fce0d3",
      "#f6bfa4",
      "#f19b72",
      "#ee7d48",
      "#ec692d",
      "#ec601e",
      "#d24f13",
      "#bb450e",
      "#a33a06"
    ],
    "orange": [
      "#fff1e8",
      "#fce0d3",
      "#f6bfa4",
      "#f19b72",
      "#ee7d48",
      "#ec692d",
      "#ec601e",
      "#d24f13",
      "#bb450e",
      "#a33a06"
    ],
    "lightOrange": [
      "#fff0e4",
      "#ffdfcf",
      "#fbbea0",
      "#f69b6d",
      "#f37c42",
      "#f16a26",
      "#f15f17",
      "#d74f0b",
      "#c04505",
      "#a83900"
    ],
    "red": [
      "#ffe9e9",
      "#ffd1d1",
      "#fba0a1",
      "#f66d6d",
      "#f34242",
      "#f12726",
      "#f11717",
      "#d7090c",
      "#c00008",
      "#a90004"
    ],
    "lightBlue": [
      "#e3f6ff",
      "#cde7ff",
      "#9bccff",
      "#65affe",
      "#3997fd",
      "#1f88fd",
      "#0a80fe",
      "#006ee3",
      "#0062cd",
      "#0054b5"
    ]
  },

  /** Index of theme.colors[color].
   *  Primary shade is used in all components to determine which color from theme.colors[color] should be used.
   *  Can be either a number (0–9) or an object to specify different color shades for light and dark color schemes.
   *  Default value `{ light: 6, dark: 8 }`
   *
   *  For example,
   *  { primaryShade: 6 } // shade 6 is used both for dark and light color schemes
   *  { primaryShade: { light: 6, dark: 7 } } // different shades for dark and light color schemes
   * */
  primaryShade: { light: 5, dark: 7 },

  /** Key of `theme.colors`, hex/rgb/hsl values are not supported.
   *  Determines which color will be used in all components by default.
   *  Default value – `blue`.
   * */
  primaryColor: "lightBlue",

  /** Function to resolve colors based on variant.
   *  Can be used to deeply customize how colors are applied to `Button`, `ActionIcon`, `ThemeIcon`
   *  and other components that use colors from theme.
   * */
  // variantColorResolver: VariantColorsResolver;

  /** Determines whether text color must be changed based on the given `color` prop in filled variant
   *  For example, if you pass `color="blue.1"` to Button component, text color will be changed to `var(--mantine-color-black)`
   *  Default value – `false`
   * */
  autoContrast: true,

  /** Determines which luminance value is used to determine if text color should be light or dark.
   *  Used only if `theme.autoContrast` is set to `true`.
   *  Default value is `0.3`
   * */
  luminanceThreshold: 0.25,

  /** font-family used in all components, system fonts by default */
  // fontFamily: string;

  /** Monospace font-family, used in code and other similar components, system fonts by default  */
  // fontFamilyMonospace: string;

  /** Controls various styles of h1-h6 elements, used in TypographyStylesProvider and Title components */
  headings: {
    fontWeight: "bold"
    // textWrap: 'wrap' | 'nowrap' | 'balance' | 'pretty' | 'stable';
    // sizes: {
    //   h1: HeadingStyle;
    //   h2: HeadingStyle;
    //   h3: HeadingStyle;
    //   h4: HeadingStyle;
    //   h5: HeadingStyle;
    //   h6: HeadingStyle;
    // };
  },

  // /** Object of values that are used to set `border-radius` in all components that support it */
  // radius: MantineRadiusValues;

  // /** Key of `theme.radius` or any valid CSS value. Default `border-radius` used by most components */
  defaultRadius: "xs",

  // /** Object of values that are used to set various CSS properties that control spacing between elements */
  // spacing: MantineSpacingValues;

  // /** Object of values that are used to control `font-size` property in all components */
  // fontSizes: MantineFontSizesValues;

  // /** Object of values that are used to control `line-height` property in `Text` component */
  // lineHeights: MantineLineHeightValues;

  // /** Object of values that are used to control breakpoints in all components,
  //  *  values are expected to be defined in em
  //  * */
  // breakpoints: MantineBreakpointsValues;

  // /** Object of values that are used to add `box-shadow` styles to components that support `shadow` prop */
  shadows: {
    "xs": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), 0 calc(0.0625rem * var(--mantine-scale)) calc(0.125rem * var(--mantine-scale)) rgba(0, 0, 0, 0.1)",
    "sm": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(0.625rem * var(--mantine-scale)) calc(0.9375rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.4375rem * var(--mantine-scale)) calc(0.4375rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale))",
    "md": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(1.25rem * var(--mantine-scale)) calc(1.5625rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.625rem * var(--mantine-scale)) calc(0.625rem * var(--mantine-scale)) calc(-0.3125rem * var(--mantine-scale))",
    "lg": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(1.75rem * var(--mantine-scale)) calc(1.4375rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(0.75rem * var(--mantine-scale)) calc(0.75rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale))",
    "xl": "0 calc(0.0625rem * var(--mantine-scale)) calc(0.1875rem * var(--mantine-scale)) rgba(0, 0, 0, 0.05), rgba(0, 0, 0, 0.05) 0 calc(2.25rem * var(--mantine-scale)) calc(1.75rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale)), rgba(0, 0, 0, 0.04) 0 calc(1.0625rem * var(--mantine-scale)) calc(1.0625rem * var(--mantine-scale)) calc(-0.4375rem * var(--mantine-scale))"
  }

  // // /** Determines whether user OS settings to reduce motion should be respected, `false` by default */
  // // respectReducedMotion: boolean;

  // // /** Determines which cursor type will be used for interactive elements
  // //  * - `default` – cursor that is used by native HTML elements, for example, `input[type="checkbox"]` has `cursor: default` styles
  // //  * - `pointer` – sets `cursor: pointer` on interactive elements that do not have these styles by default
  // //  */
  // // cursorType: 'default' | 'pointer';

  // /** Default gradient configuration for components that support `variant="gradient"` */
  // defaultGradient: MantineGradient;

  // /** Class added to the elements that have active styles, for example, `Button` and `ActionIcon` */
  // activeClassName: string;

  // /** Class added to the elements that have focus styles, for example, `Button` or `ActionIcon`.
  //  *  Overrides `theme.focusRing` property.
  //  */
  // focusClassName: string;

  // /** Allows adding `classNames`, `styles` and `defaultProps` to any component */
  // components: MantineThemeComponents;

  // /** Any other properties that you want to access with the theme objects */
  // other: MantineThemeOther;
})