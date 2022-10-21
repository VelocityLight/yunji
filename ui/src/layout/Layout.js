import {
  BarChartOutlined,
  TeamOutlined,
} from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import React, { useState } from 'react';

const { Header, Sider, Content } = Layout;

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
          items={[
            {
              key: '1',
              icon: <BarChartOutlined />,
              label: 'Realtime',
            },
            {
              key: '2',
              icon: <TeamOutlined />,
              label: 'Distribution',
            },
            // {
            //     key: '3',
            //     icon: <UploadOutlined />,
            //     label: 'nav 3',
            // },
          ]}
        />
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
