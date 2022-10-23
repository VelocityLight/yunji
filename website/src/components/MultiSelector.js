import { SelectProps, Select, Space } from 'antd';
import React from 'react';
const { Option } = Select;

export const handleChange = (onSelect) => {
  return (value) => {
    return onSelect(value)
  }
};

const MultiSelector = ({ items = [], onSelect = (newValue) => { return newValue } }) => {
  const options = [];
  for (let i = 0; i < items.length; i++) {
    const value = items[i]
    // children.push(<Option key={i.toString(items.length) + i}>{items[i]}</Option>);
    options.push({
      label: `${value}`,
      value,
    });
  }

  const selectProps = {
    mode: 'multiple',
    style: {
      width: '100%',
    },
    items,
    options,
    onChange: (newValue) => {
      onSelect(newValue)
    },
    placeholder: 'Select Item...',
    maxTagCount: 'responsive',
  };

  return (
    <Space
      direction="vertical"
      style={{
        width: '20%',
      }}
    >
      <Select {...selectProps} />
    </Space >
  );
};

export default MultiSelector;

