import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { ChartCanvas, Chart } from 'react-stockcharts';
import { fitWidth } from 'react-stockcharts/lib/helper';
import {
  BarSeries,
  AreaSeries,
  CandlestickSeries,
  LineSeries,
  MACDSeries
} from 'react-stockcharts/lib/series';
import {
  CrossHairCursor,
  CurrentCoordinate,
  MouseCoordinateX,
  MouseCoordinateY
} from 'react-stockcharts/lib/coordinates';
import {
  OHLCTooltip,
  MovingAverageTooltip,
  MACDTooltip
} from 'react-stockcharts/lib/tooltip';
import { ema, sma, macd } from 'react-stockcharts/lib/indicator';
import { discontinuousTimeScaleProvider } from 'react-stockcharts/lib/scale';
import { XAxis, YAxis } from 'react-stockcharts/lib/axes';
import { last, isDefined } from 'react-stockcharts/lib/utils';
import { format } from 'd3-format';
import { timeFormat } from 'd3-time-format';

const ema26 = ema()
  .id(0)
  .options({ windowSize: 26 })
  .merge((d, c) => {
    d.ema26 = c;
  })
  .accessor(d => d.ema26);

const ema12 = ema()
  .id(1)
  .options({ windowSize: 12 })
  .merge((d, c) => {
    d.ema12 = c;
  })
  .accessor(d => d.ema12);

const macdCalculator = macd()
  .options({
    fast: 12,
    slow: 26,
    signal: 9
  })
  .merge((d, c) => {
    d.macd = c;
  })
  .accessor(d => d.macd);

const macdAppearance = {
  stroke: {
    macd: '#cc0000',
    signal: '#00cc00'
  },
  fill: {
    divergence: '#4682B4'
  }
};

const smaVolume50 = sma()
  .id(3)
  .options({
    windowSize: 10,
    sourcePath: 'volume'
  })
  .merge((d, c) => {
    d.smaVolume50 = c;
  })
  .accessor(d => d.smaVolume50);

const candleAppearance = {
  wickStroke: d => (d.close > d.open ? '#00ff00' : '#ff0000'),
  fill: d => (d.close > d.open ? '#00cc00' : '#cc0000'),
  stroke: d => (d.close > d.open ? '#00ff00' : '#ff0000'),
  candleStrokeWidth: 0.9
};

class CandlestickChart extends Component {
  constructor(props) {
    super(props);

    const { data: initialData } = props;

    const calculatedData = macdCalculator(ema12(ema26(initialData)));
    const xScaleProvider = discontinuousTimeScaleProvider.inputDateAccessor(
      d => d.date
    );
    const { data, xScale, xAccessor, displayXAccessor } = xScaleProvider(
      calculatedData
    );

    const start = xAccessor(last(data));
    const end = xAccessor(data[Math.max(0, data.length - 150)]);
    const xExtents = [start, end];

    this.state = {
      data,
      xScale,
      xAccessor,
      displayXAccessor,
      xExtents
    };
  }

  render() {
    const { type, width, ratio } = this.props;
    const { data, xScale, xAccessor, displayXAccessor, xExtents } = this.state;

    const yExtents1 = isDefined(this.state.yExtents1)
      ? this.state.yExtents1
      : [d => [d.high, d.low], ema26.accessor(), ema12.accessor()];

    const yExtents3 = isDefined(this.state.yExtents3)
      ? this.state.yExtents3
      : macdCalculator.accessor();

    return (
      <ChartCanvas
        seriesName="stuff"
        type={type}
        data={data}
        height={700}
        width={width}
        ratio={ratio}
        margin={{ left: 70, right: 70, top: 20, bottom: 30 }}
        xScale={xScale}
        xAccessor={xAccessor}
        displayXAccessor={displayXAccessor}
        xExtents={xExtents}
      >
        <Chart
          id={1}
          height={500}
          yPanEnabled={isDefined(this.state.yExtents1)}
          yExtents={yExtents1}
          padding={{ top: 10, bottom: 20 }}
        >
          <XAxis
            axisAt="bottom"
            orient="bottom"
            showTicks={false}
            outerTickSize={0}
          />
          <YAxis axisAt="right" orient="right" ticks={10} />
          <CandlestickSeries {...candleAppearance} />
          <LineSeries yAccessor={ema26.accessor()} stroke={ema26.stroke()} />
          <LineSeries yAccessor={ema12.accessor()} stroke={ema12.stroke()} />
          <CurrentCoordinate
            yAccessor={ema26.accessor()}
            fill={ema26.stroke()}
          />
          <MouseCoordinateY
            at="right"
            orient="right"
            displayFormat={format('.2f')}
          />
          <CurrentCoordinate
            yAccessor={ema12.accessor()}
            fill={ema12.stroke()}
          />

          <OHLCTooltip origin={[-40, 0]} />
          <MovingAverageTooltip
            onClick={e => console.log(e)}
            origin={[-38, 15]}
            options={[
              {
                yAccessor: ema26.accessor(),
                type: ema26.type(),
                stroke: ema26.stroke(),
                windowSize: ema26.options().windowSize
              },
              {
                yAccessor: ema12.accessor(),
                type: ema12.type(),
                stroke: ema12.stroke(),
                windowSize: ema12.options().windowSize
              }
            ]}
          />
        </Chart>
        <Chart
          id={2}
          height={150}
          yExtents={[d => d.volume, smaVolume50.accessor()]}
          origin={(w, h) => [0, h - 300]}
        >
          <YAxis
            axisAt="left"
            orient="left"
            ticks={5}
            tickFormat={format('.2s')}
          />

          <MouseCoordinateY
            at="left"
            orient="left"
            displayFormat={format('.4s')}
          />

          <BarSeries
            yAccessor={d => d.volume}
            fill={d => (d.close > d.open ? '#00cc00' : '#cc0000')}
          />
          <AreaSeries
            yAccessor={smaVolume50.accessor()}
            stroke={smaVolume50.stroke()}
            fill={smaVolume50.fill()}
          />
        </Chart>
        <Chart
          id={3}
          height={150}
          yExtents={yExtents3}
          yPanEnabled={isDefined(this.state.yExtents3)}
          origin={(w, h) => [0, h - 150]}
          padding={{ top: 10, bottom: 10 }}
        >
          <XAxis axisAt="bottom" orient="bottom" />
          <YAxis axisAt="right" orient="right" ticks={2} />
          <MouseCoordinateX
            at="bottom"
            orient="bottom"
            displayFormat={timeFormat('%Y-%m-%d')}
          />
          <MouseCoordinateY
            at="right"
            orient="right"
            displayFormat={format('.2f')}
          />
          <MACDSeries yAccessor={d => d.macd} {...macdAppearance} />
          <MACDTooltip
            origin={[-38, 15]}
            yAccessor={d => d.macd}
            options={macdCalculator.options()}
            appearance={macdAppearance}
          />
        </Chart>
        <CrossHairCursor />
      </ChartCanvas>
    );
  }
}

CandlestickChart.propTypes = {
  data: PropTypes.array.isRequired,
  width: PropTypes.number.isRequired,
  ratio: PropTypes.number.isRequired,
  type: PropTypes.oneOf(['hybrid', 'svg']).isRequired
};

CandlestickChart.defaultProps = {
  type: 'hybrid'
};

export default fitWidth(CandlestickChart);
