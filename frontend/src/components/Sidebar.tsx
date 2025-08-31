import { Divider, Drawer, ListItem,ListItemIcon, ListItemText, Toolbar } from "@mui/material";
import ListItemButton from '@mui/material/ListItemButton';
import List from '@mui/material/List';
import { RouteConfig } from "../routes";
import { Link } from "react-router-dom";
const sidebarWidth = 240;

interface SidebarProps {
  routes: RouteConfig[];
}

interface ListItemLinkProps {
  icon?: React.ReactElement<unknown>;
  primary: string;
  to: string;
}


function ListItemLink(props: ListItemLinkProps) {
  const { icon, primary, to } = props;

  return (
    <ListItemButton component={Link} to={to}>
      {icon ? <ListItemIcon>{icon}</ListItemIcon> : null}
      <ListItemText primary={primary} />
    </ListItemButton>
  );
}


export default function Sidebar({routes}: SidebarProps){
    return (
      <Drawer
        sx={{
          width: sidebarWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: sidebarWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
            <Toolbar/> 
            <Divider/>
            <List>
                {routes.map((item)=>(
                    <ListItem key={item.primary} disablePadding>
                        <ListItemLink to={item.to} primary={item.primary} />
                    </ListItem>
                ))}
            </List>
        </Drawer>
    )
}