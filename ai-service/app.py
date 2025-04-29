import os
from flask import Flask
from routes.ai_speed_to_text_routes import summary_routes

app = Flask(__name__)
app.register_blueprint(summary_routes)

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))  
    app.run(host="0.0.0.0", port=port, debug=True)
