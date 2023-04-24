import React, { useState } from 'react';
import {Upload} from '../index'

const steps = [
  { id: 1, title: '上传文件' },
  { id: 2, title: '开始翻译文件' },
  { id: 3, title: '下载文件' },
];

const Stepper = () => {
  const [activeStep, setActiveStep] = useState(0);

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  return (
    <div>
      <h1>epub-translator</h1>
      <div>
        {steps.map((step) => (
          <div key={step.id}>
            <span>{step.title}</span>
          </div>
        ))}
      </div>
      <div>
        {activeStep === 0 && (
          <div>
            <Upload/>
            <button onClick={handleNext}>下一步</button>
          </div>
        )}
        {activeStep === 1 && (
          <div>
            <p>开始翻译文件组件</p>
            <button onClick={handleBack}>上一步</button>
            <button onClick={handleNext}>下一步</button>
          </div>
        )}
        {activeStep === 2 && (
          <div>
            <p>下载文件组件</p>
            <button onClick={handleBack}>上一步</button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Stepper;
