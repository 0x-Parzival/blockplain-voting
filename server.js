const express = require('express');
const fs = require('fs');
const path = require('path');
const app = express();

// Middleware
app.use(express.json());
app.use(express.static('public')); // Serve static files from "public"

// POST endpoint to save vote
app.post('/save-block', (req, res) => {
  const { vote, regionName } = req.body;

  // Map vote to hex color
  const colorMap = {
    A: '#ff5733',
    B: '#33aaff'
  };
  const hexCode = colorMap[vote] || '#000000';

  const entry = `Region: ${regionName}, Vote: ${vote}, Hex: ${hexCode}\n`;

  const filePath = path.join(__dirname, 'data', 'blocks.txt');

  fs.appendFile(filePath, entry, (err) => {
    if (err) {
      console.error('Error writing file:', err);
      return res.status(500).send('Failed to save vote.');
    }
    res.send('Vote saved.');
  });
});

// Start server
app.listen(3000, () => {
  console.log('Server running at http://localhost:3000');
});
