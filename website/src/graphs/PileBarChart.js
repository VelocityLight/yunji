import React from 'react';
import { Column } from '@ant-design/plots';



const PileBarChart = ({ data, xfiled_key, yfiled_key, serie_key }) => {
  const config = {
    data,
    isStack: true,
    xField: xfiled_key,
    yField: yfiled_key,
    seriesField: serie_key,
    label: {
      // 可手动配置 label 数据标签位置
      position: 'middle', // 'top', 'bottom', 'middle'
    },
    interactions: [
      {
        type: 'active-region',
        enable: false,
      },
    ],
    connectedArea: {
      style: (oldStyle, element) => {
        return {
          fill: 'rgba(0,0,0,0.25)',
          stroke: oldStyle.fill,
          lineWidth: 0.5,
        };
      },
    },
  };

  return <Column {...config} />;
};

export default PileBarChart;

