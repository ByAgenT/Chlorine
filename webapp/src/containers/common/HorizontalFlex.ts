import styled from 'styled-components';

// TODO: make types specific to flex values
interface HorizontalFlexProps {
  justifyContent?: string;
  alignItems?: string;
  width?: string;
}

const HorizontalFlex = styled.div<HorizontalFlexProps>`
  display: flex;
  flex-direction: row;
  ${({ width }) => width && `width: ${width};`}
  ${({ justifyContent }) => justifyContent && `justify-content: ${justifyContent};`}
  ${({ alignItems }) => alignItems && `align-items: ${alignItems};`}
`;

export { HorizontalFlex };
