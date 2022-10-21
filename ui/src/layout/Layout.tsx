import * as React from "react";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import { Navbar } from "./Navbar";
import { Sidebar } from "./Sidebar";

const mdTheme = createTheme();

export const Layout = (props) => {
    const { children } = props;

    const [open, setOpen] = React.useState(true);
    const toggleDrawer = () => {
        setOpen(!open);
    };

    return (
        <>
            <ThemeProvider theme={mdTheme}>
                <Box sx={{ display: "flex" }}>
                    <CssBaseline />
                    {/* <Navbar open={open} toggleDrawer={toggleDrawer} /> */}
                    <Sidebar open={open} toggleDrawer={toggleDrawer} />
                    <Box
                        component="main"
                        sx={{
                            backgroundColor: (theme) =>
                                theme.palette.mode === "light"
                                    ? theme.palette.grey[100]
                                    : theme.palette.grey[900],
                            flexGrow: 1,
                            height: "100vh",
                            overflow: "auto",
                        }}
                    >
                        {/* <Toolbar /> */}
                        {children}
                    </Box>
                </Box>
            </ThemeProvider>
        </>
    );
};

export default Layout;
