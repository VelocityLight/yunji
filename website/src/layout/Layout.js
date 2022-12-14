import {
  BarChartOutlined,
  TeamOutlined,
  FireOutlined
} from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import React, { useState } from 'react';
import { Link } from 'react-router-dom'

const { Sider, Content } = Layout;

const MyLayout = (props) => {
  const { children } = props;
  const [collapsed, setCollapsed] = useState(false);
  return (
    <Layout style={{
      minHeight: '100vh',
    }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={['1']}

        >
          <Menu.Item key="billing_trend">
            <BarChartOutlined /><Link to={"/home/trend"}>Billing</Link>
          </Menu.Item>
          <Menu.Item key="team">
            <TeamOutlined /><Link to={"/home/team"}>Team</Link>
          </Menu.Item>
          <Menu.Item key="realtime">
            <FireOutlined /><Link to={"/home/realtime"}>Realtime</Link>
          </Menu.Item>

        </Menu>
      </Sider>
      <Layout className="site-layout">
        <Content
          className="site-layout-background"
          style={{
            margin: '24px 16px',
            padding: 24,
            minHeight: 280,
          }}
        >
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};

export default MyLayout;
