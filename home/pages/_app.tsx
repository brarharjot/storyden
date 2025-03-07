import Head from "next/head";
import { DefaultSeo } from "next-seo";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import "../fonts.css";

function MyApp({ Component, pageProps }) {
  return (
    <>
      <Head>
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/apple-touch-icon.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon-16x16.png"
        />
        <link rel="manifest" href="/site.webmanifest" />
        <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#27b981" />
        <meta name="msapplication-TileColor" content="#27b981" />
        <meta name="theme-color" content="#27b981" />
      </Head>

      <ChakraProvider
        theme={extendTheme({
          fonts: {
            heading: "p22-mackinac-pro",
            body: "halyard-display",
          },
        })}
      >
        <DefaultSeo
          title="Homepage"
          titleTemplate="Storyden | %s"
          description="With a fresh new take on traditional bulletin board web forum software, Storyden is a modern, secure and extensible platform for building communities."
        />
        <Component {...pageProps} />
      </ChakraProvider>
    </>
  );
}

export default MyApp;
