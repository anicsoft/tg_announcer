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
    "dark": ["#f2f5ff", "#d4ddf7", "#b8c6f0", "#9bafe8", "#110b40", "#0c0831", "#080521", "#040312", "#000003", "#000"],
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
    "lightBlue": [
      "#e7f2ff",
      "#cee2ff",
      "#9cc2ff",
      "#66a0fe",
      "#3b83fd",
      "#2171fd",
      "#1168fe",
      "#0057e3",
      "#004dcb",
      "#0043b4"
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
  primaryShade: { light: 7, dark: 9 },

  /** Key of `theme.colors`, hex/rgb/hsl values are not supported.
   *  Determines which color will be used in all components by default.
   *  Default value – `blue`.
   * */
  primaryColor: "orange",

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
  }

  // /** Object of values that are used to set `border-radius` in all components that support it */
  // radius: MantineRadiusValues;

  // /** Key of `theme.radius` or any valid CSS value. Default `border-radius` used by most components */
  // defaultRadius: MantineRadius;

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
  // shadows: MantineShadowsValues;

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