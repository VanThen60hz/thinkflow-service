import os
from flask import Flask
from routes.ai_transcript_routes import transcript_routes
from routes.ai_summary_routes import summary_routes
from routes.ai_mindmap_routes import mindmap_routes

app = Flask(__name__)
app.register_blueprint(transcript_routes)
app.register_blueprint(summary_routes)
app.register_blueprint(mindmap_routes)

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))  
    app.run(host="0.0.0.0", port=port, debug=True)
