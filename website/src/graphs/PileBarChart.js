import React from 'react';
import { Column } from '@ant-design/plots';
import { each, groupBy } from '@antv/util';
import { useEffect, useState } from 'react';
import { useQuery } from "react-query";
import { fetchBillingByTagAndService } from "../api/billing_api"

const PileBarChart = ({ data = [], xfiled = "time", yfield = "cost", serie = "service" }) => {
  if (yfield == "cost") {
    data.map(
      item => {
        item.cost = Math.floor(item[yfield])
        return item
      }
    )
  }

  const annotations = [];
  each(groupBy(data, xfiled), (data, k) => {
    const value = data.reduce((a, b) => a + b['cost'], 0);
    annotations.push({
      type: 'text',
      position: [k, value],
      content: `${value}`,
      style: {
        textAlign: 'center',
        fontSize: 14,
        fill: 'rgba(0,0,0,0.85)',
      },
      offsetY: -10,
    });
  });

  const config = {
    data,
    isStack: true,
    xField: 'time',
    yField: 'cost',
    seriesField: 'service',
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
    xAxis: {
      label: {
        autoRotate: false,
      },
    },
    slider: {
      start: 0,
      end: 1,
    },
    annotations,
  };
  return <Column {...config} />;
};

export default PileBarChart;
