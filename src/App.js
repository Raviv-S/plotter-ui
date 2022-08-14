/*
 * Copyright 2017-2021 Elyra Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint no-console: ["error", { allow: ["log", "info", "warn", "error", "time", "timeEnd"] }] */

import React from "react";
import PropTypes from "prop-types";
import { IntlProvider } from "react-intl";
import pipeline from "./generated-pipeline.json";
import paramDef from "./generated-params.json";
import { CommonCanvas, CanvasController } from "@elyra/canvas";
import { CommonProperties } from "@elyra/canvas";
import "./index.scss";


class App extends React.Component {
	constructor(props) {
		super(props);

		this.canvasController = new CanvasController();
		this.canvasController.setPipelineFlow(pipeline);
		this.canvasController.autoLayout();
		this.canvasController.autoLayout("horizontal", this.canvasController.getSupernodes()[0].subflow_ref.pipeline_id_ref);
		
		for (let i = 0; i < this.canvasController.getSupernodes().length; i++) {
			this.canvasController.autoLayout("vertical", this.canvasController.getSupernodes()[i].subflow_ref.pipeline_id_ref)
		}
		this.state = {
			showPropertiesDialog: false,
			propertiesInfo: {
				parameterDef: paramDef,
			},
		}
	}
	
	// This function is incase we want to design each node differently (Need to uncomment the part in the return function)
	layoutHandler(data) {
		const labLen = data.label ? data.label.length : 0;
		let width = 120;
		let bodyPath = "";
		let selectionPath = "";
		let imageWidth = 80;
		let imageHeight = 80;
		// let imagePosX = 120;
		// let imagePosY = 110;

		switch (data.op) {
		case "rectangle": {
			bodyPath = "     M  0 0  L  0 60 120 60 120  0  0  0 Z";
			selectionPath = "M -5 -5 L -5 65 125 65 125 -5 -5 -5 Z";
			break;
		}
		case "pentagon": {
			bodyPath = "     M  0 20 L  0 60 120 60 120 20 60  0  0 20 Z";
			selectionPath = "M -5 17 L -5 65 125 65 125 17 60 -5 -5 17 Z";
			break;
		}
		case "octagon": {
			bodyPath = "     M  0 20 L  0 40  20 60 100 60 120 40 120 20 100 0  20  0 Z";
			selectionPath = "M -5 20 L -5 40  20 65 100 65 125 40 125 20 100 -5 20 -5 Z";
			break;
		}
		case "ellipse": {
			bodyPath = "     M  0 30 Q  0  0 60  0 Q 120  0 120 30 Q 120 60 60 60 Q  0 60  0 30 Z";
			selectionPath = "M -5 30 Q -5 -5 60 -5 Q 125 -5 125 30 Q 125 65 60 65 Q -5 65 -5 30 Z";
			break;
		}
		case "triangle": {
			bodyPath = "     M   0 60 L  140 60 70  0 0 60 Z";
			selectionPath = "M  -5 65 L  145 65 70 -5 5 65 Z";
			imageWidth = 80;
			imageHeight = 80;
			// imagePosX = 120;
			// imagePosY = 110;
			break;
		}
		case "hexagon": {
			width = (labLen * 9) + 60; // Allow 9 pixels for each character
			const corner = width - 30;
			bodyPath = `     M   0 30 L 30 60 ${corner} 60 ${width}     30 ${corner}  0 30  0 Z`;
			selectionPath = `M  -5 30 L 30 65 ${corner} 65 ${width + 5} 30 ${corner} -5 30 -5 Z`;
			break;
		}
		default:
			return {};
		}

		const nodeFormat = {
			defaultNodeWidth: width, // Override default width with calculated width
			labelPosX: (width / 2), // Specify center of label as center of node Note: text-anchor is set to middle in the CSS for this label
			labelMaxWidth: width, // Set big enough so that label is not truncated and so no ... appears
			ellipsisPosX: width - 25, // Always position 25px in from the right side
			bodyPath: bodyPath,
			selectionPath: selectionPath,
			imageWidth: imageWidth,
			imageHeight: imageHeight
		};

		return nodeFormat;
	}

	// Those functions handle the properties shown when clicking on a node.
	closePropertiesEditorDialog = this.closePropertiesEditorDialog.bind(this);
	clickActionHandler = this.clickActionHandler.bind(this);

	closePropertiesEditorDialog() {
		console.log(this.state.showPropertiesDialog)
		this.currentEditorId = null;
		this.setState({ showPropertiesDialog: false, propertiesInfo: {} });
	}

	clickActionHandler(data) {
		if (data.objectType === "node"){
			this.setState({ showPropertiesDialog: true, propertiesInfo: {parameterDef: paramDef[data.id]} });
		} else {
			this.setState({showPropertiesDialog: false})
		}
	}

	getCommonProperties() {
		const propertiesConfig = {
			containerType: "Custom",
			rightFlyout: true
		}
		const propertiesInfo = this.state.propertiesInfo
		const callbacks = {
			closePropertiesDialog: this.closePropertiesEditorDialog,
			actionHandler: this.clickActionHandler
		};

		const commonProperties = (
			<CommonProperties
				ref={(instance) => {
					this.CommonProperties = instance;
				} }
				propertiesInfo={propertiesInfo}
				propertiesConfig={propertiesConfig}
				callbacks={callbacks}
				showRightFlyout={true}
				clickActionHandler= {this.clickActionHandler}
			/>);

		return commonProperties;
	}

	
	render() {
		
		let showRightFlyoutProp = this.state.showPropertiesDialog;
		const commonCanvasConfig = {
			//enableLinkType: "Elbow",
			enableNodeFormatType: "Horizontal",
			enableRightFlyoutUnderToolbar: true,
			// enablePaletteLayout: "None",
			// enableConnectionType: "Halo",
			// enableParentClass: "nodes-colors",
			// //enableSaveZoom: "LocalStorage",
			// enableNodeLayout: {
				// 	bodyPath: "     M  0 30 Q  0  0 60  0 Q 120  0 120 30 Q 120 60 60 60 Q  0 60  0 30 Z",
				// 	selectionPath: "M -5 30 Q -5 -5 60 -5 Q 125 -5 125 30 Q 125 65 60 65 Q -5 65 -5 30 Z",
				// 	defaultNodeWidth: 120,
				// 	defaultNodeHeight: 100,
				// 	imageWidth: 30,
				// 	imageHeight: 30,
				// 	imagePosX: 20,
				// 	imagePosY: 10,
				// 	labelEditable: true,
				// 	labelPosX: 60,
				// 	labelPosY: 37,
				// 	labelWidth: 90,
				// 	labelHeight: 17, // Should match the font size specified in css + padding
				// 	ellipsisDisplay: true,
				// 	ellipsisPosX: 100,
				// 	ellipsisPosY: 20,
				// 	haloDisplay: false,
				// 	portPosY: 30
				// }
			
		}
		const rightFlyoutContent = this.getCommonProperties();
		return (
			<div id="harness-app-container">
				<IntlProvider locale="en">
					<CommonCanvas
						canvasController={this.canvasController}
						config={commonCanvasConfig}
						// layoutHandler={this.layoutHandler}
						showRightFlyout={showRightFlyoutProp}
						rightFlyoutContent={rightFlyoutContent}
						clickActionHandler= {this.clickActionHandler}
					/>
				</IntlProvider>
			</div>
		);
	}
}

App.propTypes = {
	config: PropTypes.object
};

export default App;
