import React, { useState } from 'react';

const Upload = () => {
  const [file, setFile] = useState(null);
  const [filename, setFilename] = useState('');
  const [uploadSuccess, setUploadSuccess] = useState(false);

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
      setUploadSuccess(true);
    } catch (error) {
      console.error(error);
      setUploadSuccess(false);
    }
  };

  return (
    <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
      <h1 style={{marginBottom: '20px'}}>上传文件</h1>
      <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
        <input type="file" onChange={handleFileChange} style={{marginBottom: '20px'}} />
        <button onClick={handleUpload} style={{padding: '10px 20px', backgroundColor: '#4CAF50', color: 'white', border: 'none', borderRadius: '5px'}}>上传</button>
        {uploadSuccess && <p style={{color: 'green'}}>上传成功</p>}
        {!uploadSuccess && <p style={{color: 'red'}}>未上传文件/上传失败</p>}
      </div>
    </div>
  );
};

export default Upload;

