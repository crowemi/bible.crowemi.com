import type { Metadata } from "next";
import "./globals.css";
import Footer from "./ui/footer";
import { montserrat } from "./ui/fonts";

export const metadata: Metadata = {
  title: "bible.crowemi.com",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
    <body className={`bg-white dark:bg-gray-900 ${montserrat.className}`}>
          <div className="pb-26 sm:pb-24">
            {children}
          </div>
          <Footer />
    </body>
    </html>
  );
}
