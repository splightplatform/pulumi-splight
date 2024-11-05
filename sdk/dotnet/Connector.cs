// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/connector:Connector [options] splight_connector.&lt;name&gt; &lt;connector_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/connector:Connector")]
    public partial class Connector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<Outputs.ConnectorInput>> Inputs { get; private set; } = null!;

        /// <summary>
        /// log level of the connector
        /// </summary>
        [Output("logLevel")]
        public Output<string?> LogLevel { get; private set; } = null!;

        /// <summary>
        /// instance size
        /// </summary>
        [Output("machineInstanceSize")]
        public Output<string?> MachineInstanceSize { get; private set; } = null!;

        /// <summary>
        /// the name of the connector to be created
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// id of the compute node where the connector runs
        /// </summary>
        [Output("node")]
        public Output<string?> Node { get; private set; } = null!;

        /// <summary>
        /// restart policy of the connector
        /// </summary>
        [Output("restartPolicy")]
        public Output<string?> RestartPolicy { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ConnectorTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// [NAME-VERSION] the version of the hub connector
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Connector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connector(string name, ConnectorArgs args, CustomResourceOptions? options = null)
            : base("splight:index/connector:Connector", name, args ?? new ConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connector(string name, Input<string> id, ConnectorState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/connector:Connector", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/splightplatform",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Connector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connector Get(string name, Input<string> id, ConnectorState? state = null, CustomResourceOptions? options = null)
        {
            return new Connector(name, id, state, options);
        }
    }

    public sealed class ConnectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ConnectorInputArgs>? _inputs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ConnectorInputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ConnectorInputArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// log level of the connector
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// instance size
        /// </summary>
        [Input("machineInstanceSize")]
        public Input<string>? MachineInstanceSize { get; set; }

        /// <summary>
        /// the name of the connector to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// id of the compute node where the connector runs
        /// </summary>
        [Input("node")]
        public Input<string>? Node { get; set; }

        /// <summary>
        /// restart policy of the connector
        /// </summary>
        [Input("restartPolicy")]
        public Input<string>? RestartPolicy { get; set; }

        [Input("tags")]
        private InputList<Inputs.ConnectorTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ConnectorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ConnectorTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// [NAME-VERSION] the version of the hub connector
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ConnectorArgs()
        {
        }
        public static new ConnectorArgs Empty => new ConnectorArgs();
    }

    public sealed class ConnectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ConnectorInputGetArgs>? _inputs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ConnectorInputGetArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ConnectorInputGetArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// log level of the connector
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// instance size
        /// </summary>
        [Input("machineInstanceSize")]
        public Input<string>? MachineInstanceSize { get; set; }

        /// <summary>
        /// the name of the connector to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// id of the compute node where the connector runs
        /// </summary>
        [Input("node")]
        public Input<string>? Node { get; set; }

        /// <summary>
        /// restart policy of the connector
        /// </summary>
        [Input("restartPolicy")]
        public Input<string>? RestartPolicy { get; set; }

        [Input("tags")]
        private InputList<Inputs.ConnectorTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ConnectorTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ConnectorTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// [NAME-VERSION] the version of the hub connector
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ConnectorState()
        {
        }
        public static new ConnectorState Empty => new ConnectorState();
    }
}