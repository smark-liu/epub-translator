import React, { useState } from 'react';

const Upload = () => {
  const [file, setFile] = useState(null);
  const [filename, setFilename] = useState('');

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
    setFilename(e.target.files[0].name);
  };

  const handleUpload = async () => {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('filename', filename);

    try {
      const response = await fetch('/api/upload', {
        method: 'POST',
        body: formData,
      });
      const result = await response.json();
      console.log(result);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <h1>上传文件</h1>
      <div>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleUpload}>上传</button>
      </div>
    </div>
  );
};

export default Upload;
