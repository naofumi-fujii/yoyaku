import type { Meta, StoryObj } from '@storybook/react';
import Typography from '../../components/atoms/Typography';

const meta: Meta<typeof Typography> = {
  component: Typography,
  title: 'Atoms/Typography',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Typography>;

export const Heading1: Story = {
  args: {
    variant: 'h1',
    children: 'Heading 1',
  },
};

export const Heading2: Story = {
  args: {
    variant: 'h2',
    children: 'Heading 2',
  },
};

export const Heading3: Story = {
  args: {
    variant: 'h3',
    children: 'Heading 3',
  },
};

export const Heading4: Story = {
  args: {
    variant: 'h4',
    children: 'Heading 4',
  },
};

export const Body: Story = {
  args: {
    variant: 'body',
    children: 'Body text for paragraphs and general content',
  },
};

export const Caption: Story = {
  args: {
    variant: 'caption',
    children: 'Caption text for smaller details',
  },
};

export const Label: Story = {
  args: {
    variant: 'label',
    children: 'Label for form elements',
  },
};

export const CustomColor: Story = {
  args: {
    variant: 'body',
    color: 'text-blue-500',
    children: 'Text with custom color',
  },
};